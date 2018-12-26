package worker

import "web-crawler/engine"

type CrawlerService struct {

}

var conEngine engine.ConcurrentEngine

func (CrawlerService) Process (req Request, result *ParseResult) error {
	engineReq := DeserializeRequest(req)
	parseResult, err := conEngine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(parseResult)
	return nil
}