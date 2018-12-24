package parser

import (
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	all := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range all {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser(),
		})
	}
	return result
}
