package main

import "fmt"

func main() {
	for x := 0; x < 100; x++ {
		for y := 0; y < 50; y++ {
			if x == 15 && y == 25 {
				goto breakFor
			}
		}
	}

	return

breakFor:
	fmt.Println("done")
}
