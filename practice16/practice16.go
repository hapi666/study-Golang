package main

import "fmt"

/*
	看这个代码会疑惑结果为何不是4
   调用任何一个函数前都要先对函数的参数进行求值，之后再进入函数体，只不过defer将进入函数执行的过程推迟到defer的调用方退出之前了
*/

func main() {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }())
	i++
}
