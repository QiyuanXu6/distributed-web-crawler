package worker

import (
	"web-crawler/crawler_distributed/config"
	"web-crawler/engine"
	"web-crawler/zhenai/parser"
)

type SerializedParser struct {
	FuncName string
	Args interface{}
}

// A wrapper for engine request (because the original request contains Parser(interface contains function(SeriParser(name)) and args)
type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request // Serialized requests
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			FuncName: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) engine.Request {
	funcName := r.Parser.FuncName
	return engine.Request{
		Url: r.Url,
		Parser: func(args interface{}) engine.Parser {
			switch funcName {
			case config.ParseCityList:
				return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)
			case config.ParseCity:
				return engine.NewFuncParser(parser.ParseCity, config.ParseCity)
			case config.ParseProfile:
				if userName, ok := args.(string); ok {
					return parser.NewProfileParser(userName)
				} else {
					panic("The username is profile parser is not a string")
				}
			default:
				panic("Unexpected parser funcName")
			}
		}(r.Parser.Args),
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		funcName := req.Parser.FuncName
		result.Requests = append(result.Requests, engine.Request{
			Url: req.Url,
			Parser: func(args interface{}) engine.Parser {
				switch funcName {
				case config.ParseCityList:
					return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)
				case config.ParseCity:
					return engine.NewFuncParser(parser.ParseCity, config.ParseCity)
				case config.ParseProfile:
					if userName, ok := args.(string); ok {
						return parser.NewProfileParser(userName)
					} else {
						panic("The username is profile parser is not a string")
					}
				default:
					panic("Unexpected parser funcName")
				}
			}(req.Parser.Args),
		})
	}
	return result
}


