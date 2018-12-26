package main

import (
	"web-crawler/crawler_distributed/persist/client"
	"web-crawler/engine"
	"web-crawler/scheduler"
	"web-crawler/zhenai/parser"
)

func main() {
	itemchan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 6,
		DedupService: *engine.NewDedupService(),
		ItemChan: itemchan,
	}
	//e.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
