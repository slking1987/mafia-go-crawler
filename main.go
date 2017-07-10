package main

import (
	"time"

	"mafia-go-crawler/common/log"
	"mafia-go-crawler/model"
	"mafia-go-crawler/module/channel"
	"mafia-go-crawler/module/crawler"
	"mafia-go-crawler/module/processor"
	"mafia-go-crawler/module/output"
	"mafia-go-crawler/module/monitor"
)

func dummyInput() {
	log.Info("load dummy input url.....")
	tempUrl := &model.Input{Url: "http://slking1987.com"}
	channel.GLOBAL_CHAN_MANAGER.PushInput(tempUrl)
}

func main() {
	// 加载待爬取任务列表
	go dummyInput()
	// 启动爬取
	time.Sleep(2 * time.Second)
	crawler.GLOBAL_CRAWLER_MANAGER.Start()
	// 启动处理
	time.Sleep(2 * time.Second)
	processor.GLOBAL_PROCESSOR_MANAGER.Start()
	// 启动输出
	time.Sleep(2 * time.Second)
	go output.GLOBAL_OUTPUT_MANAGER.Start()

	// 启动监控 定期输出系统运行状况
	go monitor.GLOBAL_MONITOR_MANAGER.Start()

	// 设置main进程运行时间
	select {
	case <-time.After(30 * time.Minute):
		log.Error("it is time to stop...")
	}
}
