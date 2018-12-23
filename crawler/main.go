package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.zhenai.com/zhenghun", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)

	//resp, err := http.Get("https://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		fmt.Println("status code wrong")
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println("%s\n", string(all))
	printCityList(all)
}


func printCityList(contents []byte) {
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`)
	all := reg.FindAllSubmatch(contents, -1)
	for _, m := range all {
		for _, c := range m {
			fmt.Println(string(c))
		}
	}
	fmt.Println(len(all))
}
