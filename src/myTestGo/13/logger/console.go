package main

import (
	"fmt"
	"os"
)

//命令行写入器
type consoleWriter struct {
}

//实现Write方法
func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))

	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
