package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "pocess mode")

func main() {
	// 解析
	flag.Parse()
	// 输出
	fmt.Println(*mode)
}
