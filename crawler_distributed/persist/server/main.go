package main

import (
	"gopkg.in/olivere/elastic.v5"
	"web-crawler/crawler_distributed/persist"
	"web-crawler/crawler_distributed/rpcsupport"
)

// Start the RPC server
func main() {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(nil)
	}
	err = rpcsupport.ServeRpc(":1234", persist.ItemSaverService{
		Client: client,
	})
	if err != nil {
		panic(err)
	}

}
