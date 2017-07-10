package output

import (
	"mafia-go-crawler/common/log"
	"mafia-go-crawler/common"
	"fmt"
	"mafia-go-crawler/module/channel"
)

var GLOBAL_OUTPUT_MANAGER Manager = Manager{make(map[string]MafiaOutput, common.MAX_OUTPUT_NUM)}

type Manager struct {
	itemMap map[string]MafiaOutput
}

func (m *Manager) add(o MafiaOutput) {
	m.itemMap[o.Name()] = o
}

func (m *Manager) Start() {
	log.Info("[Output]start....")
	for {
		for _, v := range GLOBAL_OUTPUT_MANAGER.itemMap {
			input := channel.GLOBAL_CHAN_MANAGER.GetOutput()

			err := v.Do(input)
			if err != nil {
				log.Error(fmt.Sprintf("[Output][%s]do error", v.Name()), err)
				continue
			}
		}
	}

	for i := 0; i < common.MAX_OUTPUT_NUM; i++ {
		log.Info(fmt.Sprintf("start crawler %d", i))
		go m.startSingle()
	}
}

func (m *Manager) startSingle() {
	for {
		input := channel.GLOBAL_CHAN_MANAGER.GetOutput()
		for _, v := range GLOBAL_OUTPUT_MANAGER.itemMap {
			err := v.Do(input)
			if err != nil {
				log.Error(fmt.Sprintf("[Output][%s]do error", v.Name()), err)
				continue
			}
		}
	}
}

func init() {
	GLOBAL_OUTPUT_MANAGER.add(constructConsole())
}
