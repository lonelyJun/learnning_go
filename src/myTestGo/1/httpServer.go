package main

// 引包，之后为括号，分割通过回车，不用逗号或分号
import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) //此处http对应 net/http中斜杠后的http？

	http.ListenAndServe(":8080", nil)
}
