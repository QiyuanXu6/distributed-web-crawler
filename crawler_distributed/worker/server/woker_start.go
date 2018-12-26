package main

import (
	"log"
	"web-crawler/crawler_distributed/rpcsupport"
	"web-crawler/crawler_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(":9000", worker.CrawlerService{}))
}


