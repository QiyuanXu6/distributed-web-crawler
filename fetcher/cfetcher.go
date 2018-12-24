package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get wrong status code")
	}

	return ioutil.ReadAll(resp.Body)
}

//func determineEncoding(r io.Reader) encoding.Encoder {
//	bytes, err := bufio.NewReader(r).Peek(512)
//	if err != nil {
//		log.Printf("cFetch error")
//		return unicode.UTF8
//	}
//}