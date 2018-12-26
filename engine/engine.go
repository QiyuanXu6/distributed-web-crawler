package engine

import (
	"log"
	"web-crawler/fetcher"
)

type SimpleEngine struct {

}

func (engine *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := engine.worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Print("Items %v", item)
		}
	}
}

func (SimpleEngine) worker(request Request) (ParseResult, error) {
	log.Printf("Fetching %s", request.Url)
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("error fetch in url %s: %v", request.Url, err)
		return ParseResult{}, err
	}
	return request.Parser.Parse(body, request.Url), nil
}