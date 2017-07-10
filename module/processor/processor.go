package processor

import "mafia-go-crawler/model"

type MafiaProcessor interface {
	Do(input *model.CrawlResult) (*model.ProcessResult, error)
	Name() string
}
