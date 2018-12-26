package main

import (
	"fmt"
	"testing"
	"web-crawler/crawler_distributed/config"
	"web-crawler/crawler_distributed/rpcsupport"
	"web-crawler/crawler_distributed/worker"
)

func TestCrawler(t *testing.T) {
	const addr = ":9000"
	go rpcsupport.ServeRpc(addr, worker.CrawlerService{})

	client, err := rpcsupport.NewClient(addr)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1426975040",
		Parser: worker.SerializedParser{
			FuncName: config.ParseProfile,
			Args: "不想",
		},
	}
	var result worker.ParseResult
	err = client.Call("CrawlerService.Process", req, &result)
	if err != nil {
		t.Errorf("Failed %v", err)
	} else {
		fmt.Printf("got %v", result)
	}

}
