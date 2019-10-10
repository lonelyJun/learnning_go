package main

//声明日志写入器接口
type LogWriter interface {
	//接口抽象一个写方法，写入数据，返回错误
	Write(data interface{}) error
}

//创建一个日志器
type Logger struct {
	//写日志列表，每次写则记入
	writerList []LogWriter
}

//注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	// 对于日志器，注册日志写入器（代码说法即将一个日志写入器加入写入器列表）
	l.writerList = append(l.writerList, writer)
}

//将数据写到日志中（猜测实现LogWriter的Write方法）
func (l *Logger) Log(data interface{}) {
	//遍历了日志中的所有日志写入器，将每个写入器中的date写入
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

//创建实例对象
func NewLogger() *Logger {
	return &Logger{}
}
