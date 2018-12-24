package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have size %d, but have %d size", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url is %s, but got %s", url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Errorf("expected url is %s, but got %s", city, result.Items[i])
		}
	}
}
