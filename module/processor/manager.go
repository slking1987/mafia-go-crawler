package processor

import (
	"mafia-go-crawler/common"
	"fmt"
	"mafia-go-crawler/common/log"
	"mafia-go-crawler/module/channel"
	"mafia-go-crawler/common/util"
	"time"
	"sync"
)

var GLOBAL_PROCESSOR_MANAGER Manager = Manager{
	make(map[string]MafiaProcessor, common.MAX_PROCESSOR_NUM),
	make(map[string]bool),
}

var lock sync.RWMutex

type Manager struct {
	itemMap     map[string]MafiaProcessor
	inputUrlMap map[string]bool
}

func (m *Manager) add(p MafiaProcessor) {
	m.itemMap[p.Name()] = p
}

func (m *Manager) Start() {
	for i := 0; i < common.MAX_PROCESSOR_NUM; i++ {
		log.Info(fmt.Sprintf("start processor %d", i))
		go m.startSingle()
	}
}

func (m *Manager) startSingle() {
	for {
		input := channel.GLOBAL_CHAN_MANAGER.GetProcess()
		for _, v := range GLOBAL_PROCESSOR_MANAGER.itemMap {
			result, err := v.Do(input)
			if err != nil {
				log.Error(fmt.Sprintf("[Processor][%s]do error", v.Name()), err)
				continue
			}

			// push image to output channel
			// if len(result.Images) > 0 {
				channel.GLOBAL_CHAN_MANAGER.PushOutput(result)
			// }
			// push sub url to input channel
			subUrls := result.SubUrls
			for i := 0; i < len(subUrls); i ++ {
				// duplicate check
				key := util.Sha1(subUrls[i].Url)
				lock.RLock()
				_, isDuplicate := m.inputUrlMap[key]
				lock.RUnlock()
				if isDuplicate {
					continue
				} else {
					lock.Lock()
					m.inputUrlMap[key] = true
					lock.Unlock()
				}
				channel.GLOBAL_CHAN_MANAGER.PushInput(&subUrls[i])
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func init() {
	GLOBAL_PROCESSOR_MANAGER.add(constructUrl())
	GLOBAL_PROCESSOR_MANAGER.add(constructImage())
}
