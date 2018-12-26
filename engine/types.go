package engine


type ParseResult struct {
	Requests []Request
	Items []Item
}

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

//
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialized() (name string, args interface{})
}

//type SerializedParser struct {
//	FunctionName string
//	Args interface{}
//}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url string
	Type string // elasticsearch type
	Id string
	Payload interface{}
}

