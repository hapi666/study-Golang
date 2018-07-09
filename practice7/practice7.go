package main

import (
	"fmt"
)

func main() {
	var i int
	i = 1
	func(i int) {
		i = 3
		fmt.Println(i)
	}(i)
	fmt.Println(i)
}
