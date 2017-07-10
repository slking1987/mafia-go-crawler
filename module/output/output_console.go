package output

import (
	"mafia-go-crawler/model"
	"mafia-go/common/log"
	"fmt"
	"mafia-go-crawler/common"
)

type ConsoleOutput struct {
	name string
}

func constructConsole() MafiaOutput {
	return &ConsoleOutput{name: common.OUTPUT_CONSOLE}
}

func (o *ConsoleOutput) Name() string {
	return o.name
}

func (o *ConsoleOutput) Do(input *model.ProcessResult) error {
	// TODO
	images := input.Images
	for i := 0; i < len(images); i ++ {
		log.Info(fmt.Sprintf("[Output][%s]%s,%s", o.Name(), images[i].Url, images[i].Desc))
	}

	log.Info(fmt.Sprintf("[Output][%s]%s,%s", o.Name(), input.Url.Url, input.Url.Desc))

	return nil
}
