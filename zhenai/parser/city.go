package parser

import (
	"fmt"
	"regexp"
	"web-crawler/crawler_distributed/config"
	"web-crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<tr><th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(content []byte, url string) engine.ParseResult {
	all := profileRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	fmt.Println("city list: ", len(all))
	limit := 10
	for _, m := range all {
		name := string(m[2])
		//result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: NewProfileParser(name),
		})
		limit--
		if limit < 0 {
			break
		}
	}

	submatch := cityUrlRe.FindAllSubmatch(content, -1)
	for _, m := range submatch {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}
	return result
}
