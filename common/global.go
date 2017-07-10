package common

import (
	"errors"
)

const (
	CHAN_INPUT_NAME       = "InputChan"
	CHAN_INPUT_CAPACITY   = 5000
	CHAN_PROCESS_NAME     = "ProcessChan"
	CHAN_PROCESS_CAPACITY = 5000
	CHAN_OUTPUT_NAME      = "OutputChan"
	CHAN_OUTPUT_CAPACITY  = 5000

	PROCESSOR_URL   = "processor_url"
	PROCESSOR_IMAGE = "processor_image"

	OUTPUT_CONSOLE = "output_console"

	MAX_CHAN_NUM = 8

	MAX_PROCESSOR_NUM = 5
	MAX_CRAWLER_NUM   = 5
	MAX_OUTPUT_NUM    = 5
)

var ERROR_CHAN_FULL = errors.New("通道已满")
