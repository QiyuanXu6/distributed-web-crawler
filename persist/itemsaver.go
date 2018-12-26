package persist

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"web-crawler/engine"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("Saved item %d %v", count, item)
			count++
			_, err := Save(item, client)
			if err != nil {
				log.Println("Error in saving %d %v", count, item)
			}

		}
	}()
	return out, nil
}

func Save(item engine.Item, client *elastic.Client) (id string, err error) {



	if item.Type == "" {
		return "", errors.New("must have Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	response, err := indexService.
		Do(context.Background())

	if err != nil {
		return "", err
	}
	return response.Id, nil
}