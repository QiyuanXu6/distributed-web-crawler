package main

import (
	"flag"
	"fmt"
	"log"
	"web-crawler/crawler_distributed/rpcsupport"
	"web-crawler/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "port for this worker to listen")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must input port number")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlerService{}))
}
