package channel

import (
	"mafia-go-crawler/common"
	"mafia-go-crawler/model"
	"fmt"
	"sync/atomic"
)

var GLOBAL_CHAN_MANAGER Manager = Manager{make(map[string]*MafiaChan, common.MAX_CHAN_NUM)}

type Manager struct {
	chanMap map[string]*MafiaChan
}

func (m *Manager) push(data interface{}, chanName string) {
	tempChan, _ := m.chanMap[chanName]
	err := tempChan.Push(data)
	tempChan.countIn = atomic.AddInt32(&tempChan.countIn, 1)
	if err != nil {
		tempChan.countErr = atomic.AddInt32(&tempChan.countErr, 1)
	}
}

func (m *Manager) get(chanName string) interface{} {
	tempChan, _ := m.chanMap[chanName]
	tempChan.countOut = atomic.AddInt32(&tempChan.countOut, 1)
	return tempChan.Get()
}

func (m *Manager) addChan(name string, capacity int) {
	tempChan := make(chan interface{}, capacity)
	mafiaChan := &MafiaChan{capacity: capacity, c: tempChan}
	m.chanMap[name] = mafiaChan
}

func (m *Manager) PushInput(data *model.Input) {
	m.push(data, common.CHAN_INPUT_NAME)
}
func (m *Manager) PushProcess(data *model.CrawlResult) {
	m.push(data, common.CHAN_PROCESS_NAME)
}
func (m *Manager) PushOutput(data *model.ProcessResult) {
	m.push(data, common.CHAN_OUTPUT_NAME)
}

func (m *Manager) GetInput() *model.Input {
	return m.get(common.CHAN_INPUT_NAME).(*model.Input)
}
func (m *Manager) GetProcess() *model.CrawlResult {
	return m.get(common.CHAN_PROCESS_NAME).(*model.CrawlResult)
}
func (m *Manager) GetOutput() *model.ProcessResult {
	return m.get(common.CHAN_OUTPUT_NAME).(*model.ProcessResult)
}

func (m *Manager) Stat() string {
	var chanStatArr []string
	for k, v := range m.chanMap {
		chanStatArr = append(chanStatArr, fmt.Sprintf("%s->cur:%d,in:%d,out:%d,err:%d", k, len(v.c), v.countIn, v.countOut, v.countErr))
	}
	return fmt.Sprintf("ChannelStat:{%v}", chanStatArr)
}

func init() {
	GLOBAL_CHAN_MANAGER.addChan(common.CHAN_INPUT_NAME, common.CHAN_INPUT_CAPACITY)
	GLOBAL_CHAN_MANAGER.addChan(common.CHAN_PROCESS_NAME, common.CHAN_PROCESS_CAPACITY)
	GLOBAL_CHAN_MANAGER.addChan(common.CHAN_OUTPUT_NAME, common.CHAN_OUTPUT_CAPACITY)
}
