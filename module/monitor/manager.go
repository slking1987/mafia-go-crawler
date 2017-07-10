package monitor

import (
	"time"
	"mafia-go-crawler/common/log"
	"mafia-go-crawler/module/channel"
)

type Manager struct {
}

var GLOBAL_MONITOR_MANAGER Manager = Manager{}

func (m *Manager) Start() {
	for {
		// channel stat
		log.Info(channel.GLOBAL_CHAN_MANAGER.Stat())
		time.Sleep(5 * time.Second)
	}
}
