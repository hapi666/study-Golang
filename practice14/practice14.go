package main

import "fmt"

func main() {
	/*
	无参闭包
	执行函数体是用外部的引用
	*/
	var i int = 1

	defer func() {
		fmt.Println("result0 =>", func() int { return i * 2 }())//用的是i的引用
	}()
	i++
	/*
	有参闭包
	先计算此时函数的参数列表，也是引用，但是 是此时的引用
	执行函数体时用直接用当时计算出来的参数
	*/
	var ii int = 1
	defer func(ii int){//此时计算出ii得1
		fmt.Println("result1 =>", func() int { return ii * 2 }())//用的是此时外部的ii的值，与defer fmt.Println("result =>",func() int { return i * 2 }())等价
	}(ii)
	ii++
}
