package main

import (
	"fmt"
)

func test() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := test()
	fmt.Println("1:", f())
	fmt.Println("2:", f())
	fmt.Println("3:", f())
	fmt.Println(f)
}
