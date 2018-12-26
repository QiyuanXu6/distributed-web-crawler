package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
	"web-crawler/crawler_distributed/config"
	"web-crawler/crawler_distributed/persist/client"
	"web-crawler/crawler_distributed/rpcsupport"
	client2 "web-crawler/crawler_distributed/worker/client"
	"web-crawler/engine"
	"web-crawler/scheduler"
	"web-crawler/zhenai/parser"
)

var (
	workerHosts = flag.String("worker_hosts", "", "worker hosts comma seperated")
)

func main() {
	flag.Parse()
	itemchan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	fmt.Println(strings.Split(*workerHosts, ","))
	processor, err := client2.CreateProccesor(pool)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      6,
		DedupService:     *engine.NewDedupService(),
		ItemChan:         itemchan,
		RequestProcessor: processor,
	}
	//e.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		c, err := rpcsupport.NewClient(":" + h)
		if err != nil {
			log.Printf("Error in create clients for workers %v\n", err)
		} else {
			clients = append(clients, c)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
