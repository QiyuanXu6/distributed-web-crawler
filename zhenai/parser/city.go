package parser

import (
	"fmt"
	"regexp"
	"web-crawler/engine"
)


const cityre = `<tr><th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr>`

func ParseCity(content []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityre)
	all := reg.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	fmt.Println("city list: ", len(all))
	limit := 10
	for _, m := range all {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
		limit--
		if limit < 0 {
			break
		}
	}
	return result
}
