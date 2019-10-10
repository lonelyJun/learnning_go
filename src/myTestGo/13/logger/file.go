package main

import (
	"errors"
	"fmt"
	"os"
)

//文件写入器
type fileWriter struct {
	file *os.File
}

func (f *fileWriter) SetFile(filename string) (err error) {
	//如果文件已经打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}
	//创建文件，保存文件句柄
	f.file, err = os.Create(filename)

	//如果出错，返回错误
	return err
}

//实现LogWriter的Write()方法
func (f *fileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}

	str := fmt.Sprintf("%v\n", data)

	//将数据以字节数组写入文件
	_, err := f.file.Write([]byte(str))

	return err
}

func newFileWriter() *fileWriter {
	return &fileWriter{}
}
