package main

import (
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


