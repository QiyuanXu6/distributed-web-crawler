package main

import (
	"web-crawler/engine"
	"web-crawler/scheduler"
	"web-crawler/zhenai/parser"
)

func main() {
	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 6,
	}
	e.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}


