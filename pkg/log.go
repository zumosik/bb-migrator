package pkg

import "fmt"

var L = NewLog()

func NewLog() *Log {
	return &Log{}
}

type Log struct {
}

func (l *Log) Printf(format string, v ...interface{}) {
	fmt.Printf(string(format+"\n"), v...)
}
