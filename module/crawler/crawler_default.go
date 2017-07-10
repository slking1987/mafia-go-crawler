package crawler

import (
	"fmt"
	"mafia-go-crawler/model"
	"mafia-go-crawler/common/log"
	"mafia-go-crawler/common/util"
)

type DefaultCrawler struct {
	name string
}

func constructDefault(name string) MafiaCrawler {
	return &DefaultCrawler{name: name}
}

func (c *DefaultCrawler) Name() string {
	return c.name
}

func (c *DefaultCrawler) Do(input *model.Input) (*model.CrawlResult, error) {
	log.Debug(fmt.Sprintf("[Crawler][%s]%s", c.Name(), input.Url))

	// http get page info
	pageInfo, err := util.HttpGet(input.Url)
	if err != nil {
		return nil, err
	}

	return &model.CrawlResult{Url: *input, Page: pageInfo}, nil
}
