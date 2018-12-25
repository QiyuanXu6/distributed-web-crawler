package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("Saved item %d %v", count, item)
			count++
			_, err := save(item)
			if err != nil {
				log.Println("Error in saving %d %v", count, item)
			}

		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	response, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return response.Id, nil
}