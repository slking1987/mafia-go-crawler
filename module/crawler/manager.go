package crawler

import (
	"mafia-go-crawler/common"
	"mafia-go-crawler/module/channel"
	"mafia-go-crawler/common/log"
	"fmt"
)

var GLOBAL_CRAWLER_MANAGER Manager = Manager{
	make(map[string]MafiaCrawler, common.MAX_CRAWLER_NUM),
}

type Manager struct {
	itemMap map[string]MafiaCrawler
}

func (m *Manager) add(c MafiaCrawler) {
	m.itemMap[c.Name()] = c
}

func (m *Manager) Start() {
	for i := 0; i < common.MAX_CRAWLER_NUM; i++ {
		log.Info(fmt.Sprintf("start crawler %d", i))
		go m.startSingle()
	}
}

func (m *Manager) startSingle() {
	for {
		input := channel.GLOBAL_CHAN_MANAGER.GetInput()
		for _, v := range GLOBAL_CRAWLER_MANAGER.itemMap {
			result, err := v.Do(input)
			if err != nil {
				log.Error(fmt.Sprintf("[Crawler][%s]do error", v.Name()), err)
				continue
			}
			channel.GLOBAL_CHAN_MANAGER.PushProcess(result)
		}
	}
}

func init() {
	tempCrawler := constructDefault("crawler-default")
	GLOBAL_CRAWLER_MANAGER.add(tempCrawler)
}
