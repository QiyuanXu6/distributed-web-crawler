package parser

import (
	"regexp"
	"web-crawler/crawler_distributed/config"
	"web-crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

// Function extract inforamtion and url from web contents
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	all := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 2

	for _, m := range all {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
