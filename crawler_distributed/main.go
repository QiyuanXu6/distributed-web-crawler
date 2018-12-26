package main

import (
	"web-crawler/crawler_distributed/config"
	"web-crawler/crawler_distributed/persist/client"
	client2 "web-crawler/crawler_distributed/worker/client"
	"web-crawler/engine"
	"web-crawler/scheduler"
	"web-crawler/zhenai/parser"
)

func main() {
	itemchan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	processor, err := client2.CreateProccesor()
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 6,
		DedupService: *engine.NewDedupService(),
		ItemChan: itemchan,
		RequestProcessor: processor,
	}
	//e.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}
