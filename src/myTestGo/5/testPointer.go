package main

import "fmt"

func main() {
	str := "test string"

	ptr := &str

	fmt.Printf("%p %s\n", ptr, *ptr)
}
