package parser

import (
	"regexp"
	"web-crawler/engine"
)


const cityre = `<tr><th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr>`

func ParseCity(content []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityre)
	all := reg.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	limit := 5
	for _, m := range all {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, string(m[2]))
			},
		})
		limit--
		if limit < 0 {
			break
		}
	}
	return result
}
