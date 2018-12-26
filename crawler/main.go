package main

import (
	"web-crawler/crawler_distributed/config"
	"web-crawler/engine"
	"web-crawler/persist"
	"web-crawler/scheduler"
	"web-crawler/zhenai/parser"
)

func main() {
	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	// Why I need & and * sign here?

	itemchan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	tmp := engine.ConcurrentEngine{}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 6,
		DedupService: *engine.NewDedupService(),
		ItemChan: itemchan,
		RequestProcessor: tmp.Worker,
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


