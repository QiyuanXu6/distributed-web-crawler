package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"testing"
	"web-crawler/model"
)

func TestSaver(t *testing.T) {
	profile := model.Profile{
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
	}
	id, err := save(profile)
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
		Type("zhenai").
		Id(id).
		Do(context.Background())
	log.Printf("%s", result.Source)


	var actual model.Profile
	bytes, err := result.Source.MarshalJSON()
	json.Unmarshal(bytes, &actual)

	if actual != profile {
		t.Errorf("Got %v, expected %v", actual, profile)
	}
}