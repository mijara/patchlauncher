package adapter

import "fmt"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l Logger) Info(s string) {
	fmt.Println(s)
}

func (l Logger) Debug(s string) {
	fmt.Println(s)
}
