package engine

import (
	"fmt"
	"log"
	"web-crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	DedupService DedupService
	ItemChan chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	WorkerChan() chan Request // Get In channel from Scheduler
	ReadyNotifier
	Run()
}

// Ducking typing: don't need to implement this interface
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		if e.DedupService.isDup(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for  {
		// Get one result from out chan
		result := <-out
		// Print out the result.items
		for _, item := range result.Items {
			fmt.Printf("Got item %v\n", item)
			go func() {
				e.ItemChan <- item
			}()
		}

		// Submit all result.requests to scheduler
		for _, request := range result.Requests {
			if e.DedupService.isDup(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
		fmt.Printf("finshed submit\n")
	}

}

// Create one goroutine: contains one worker, get from in chan, output into out chan
// Channels the only thing about the worker
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// work count => how many instance of this goroutine
	go func() {
		for {
			// tell scheduler I am ready to get new requests from scheduler
			ready.WorkerReady(in)

			request := <-in

			// Call RPC worker
			result, err := e.RequestProcessor(request)

			//result, err := engine.Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func (ConcurrentEngine) Worker(request Request) (ParseResult, error) {
	log.Printf("Fetching %s", request.Url)
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("error fetch in url %s: %v", request.Url, err)
		return ParseResult{}, err
	}
	log.Printf("Fetching %s succeed", request.Url)
	return request.Parser.Parse(body, request.Url), nil
}