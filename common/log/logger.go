package log

import (
	"fmt"
	"time"
)

func Debug(a ...interface{}) {
	fmt.Println(fmt.Sprintf("[MAFIA-GO][DEBUG][%s]", time.Now().Format("2006-01-02 15:04:05")), a)
}

func Info(a ...interface{}) {
	fmt.Println(fmt.Sprintf("[MAFIA-GO][INFO][%s]", time.Now().Format("2006-01-02 15:04:05")), a)
}

func Error(a ...interface{}) {
	fmt.Println(fmt.Sprintf("[MAFIA-GO][ERROR][%s]", time.Now().Format("2006-01-02 15:04:05")), a)
}

func Warn(a ...interface{}) {
	fmt.Println(fmt.Sprintf("[MAFIA-GO][WARN][%s]", time.Now().Format("2006-01-02 15:04:05")), a)
}
