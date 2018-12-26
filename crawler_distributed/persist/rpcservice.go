package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"web-crawler/engine"
	"web-crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}


func (s ItemSaverService) Save(item engine.Item, result *string) error {
	_, err := persist.Save(item, s.Client)
	if err == nil {
		*result = "OK"
	}
	return err
}
