package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"testing"
	"web-crawler/engine"
	"web-crawler/model"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url: "http://www.test.com/url",
		Type: "zhenai",
		Id: "123456",
		Payload: model.Profile{
			Name: "a",
			Age: "30",
			AvatarUrl: "http://www.test.com/avatar",
			BasicInfo: "ok",
			DetailInfo: "ok",
			Education: "本科",
			Gender: "女",
			Height: "10",
			Salary: "500000",
			Marriage: "no",
		},
	}
	_, err := Save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	result, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	log.Printf("%s", result.Source)


	var actual engine.Item
	bytes, err := result.Source.MarshalJSON()
	json.Unmarshal(bytes, &actual)
	//Unmarshal translate the interface payload into a map rather than a struct

	actualProfile, err := model.FromJsonObj(actual.Payload)

	if actualProfile != expected.Payload {
		t.Errorf("Got %v, expected %v", actualProfile, expected.Payload)
	}
}