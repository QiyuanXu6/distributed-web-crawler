package main

import (
	"web-crawler/engine"
	"web-crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}


