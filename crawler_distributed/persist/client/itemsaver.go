package client

import (
	"log"
	"web-crawler/crawler_distributed/rpcsupport"
	"web-crawler/engine"
)

// Just like ItemSaver, just return a channel
func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("Saved item %d %v\n", count, item)
			count++
			var result string
			err := client.Call("ItemSaverService.Save", item, &result)
			//_, err := Save(item, client)
			if err != nil {
				log.Printf("Error in saving %d %v %v\n", count, item, err)
			}

		}
	}()
	return out, nil
}
