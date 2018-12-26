package client

import (
	"web-crawler/crawler_distributed/rpcsupport"
	"web-crawler/crawler_distributed/worker"
	"web-crawler/engine"
)

func CreateProccesor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(":9000")
	if err != nil {
		return nil, err
	}
	return func(request engine.Request) (result engine.ParseResult, e error) {
		serializedRequest := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		err := client.Call("CrawlerService.Process", serializedRequest, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}