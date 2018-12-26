package engine


type ParseResult struct {
	Requests []Request
	Items []Item
}

type Request struct {
	Url string
	Parser Parser
}

// Defind a interface which can store some parser function, and informaiton about the parser function
// Can parse sth, can be serialized into names and args
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//type SerializedParser struct {
//	FunctionName string
//	Args interface{}
//}

// real function type to parse a page of data
// first element in Parser
type ParserFunc func(contents []byte, url string) ParseResult

// NilParser is a Parser
type NilParser struct {

}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// FuncParser is a Parser
type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name: name,
	}
}


type Item struct {
	Url string
	Type string // elasticsearch type
	Id string
	Payload interface{}
}

