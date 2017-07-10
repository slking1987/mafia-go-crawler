package channel

import (
	"mafia-go-crawler/common"
)

type MafiaChan struct {
	c        chan interface{}
	capacity int
	countIn  int32
	countOut int32
	countErr int32
}

func (c *MafiaChan) Push(data interface{}) error {
	if len(c.c) >= c.capacity {
		return common.ERROR_CHAN_FULL
	}

	c.c <- data
	return nil
}

func (c *MafiaChan) Get() interface{} {
	return <-c.c
}
