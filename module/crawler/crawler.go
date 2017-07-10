package crawler

import "mafia-go-crawler/model"

type MafiaCrawler interface {
	Do(input *model.Input) (*model.CrawlResult, error)
	Name() string
}
