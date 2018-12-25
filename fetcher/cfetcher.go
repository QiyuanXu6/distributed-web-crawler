package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimit = time.Tick(100 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	// rate limit, block when no tick
	<- rateLimit

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Mobile Safari/537.36")


	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get wrong status code: %d", resp.StatusCode)
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