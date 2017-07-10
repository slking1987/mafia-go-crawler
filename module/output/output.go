package output

import "mafia-go-crawler/model"

type MafiaOutput interface {
	Do(input *model.ProcessResult) error
	Name() string
}
