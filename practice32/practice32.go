package main

import "fmt"

func c() (i int) {
	return 1
}

func Print1(p []int) {
	fmt.Printf(" 切片-p: %p 切片-v: %v 切片长度：%v 切片容量：%v \n", p, p, len(p), cap(p))
}

func main() {
	fmt.Println(c())
	fmt.Println("--------当切片容量是10， 长度都是10的情况-------")
	sliceRunV3 := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		sliceRunV3 = append(sliceRunV3, i)
		Print1(sliceRunV3)
	}
	// 结论： 切片长度动态变量，容量是长度加起初容量，内存不从新分配，效率高，会出现默认值，不利于操作。

	fmt.Println("--------当切片容量是0， 长度都是0的情况, 指针用法-------")
	sliceRunV4 := make([]int, 0)
	slicepre := &sliceRunV4
	for i := 0; i < 10; i++ {
		sliceRunV4 = append(sliceRunV4, i)
		fmt.Printf("切片-p:%p \t指针-p:%p \t指针的值-v:%v\t 指针值的长度: %d \t指针值的容量：%d \n", sliceRunV4, slicepre, *slicepre, len(*slicepre), cap(*slicepre))
	}
	// 结论： 指针的效果和切片容量和长度为0时，差不多，指针的长度和容量是动态变化的。
}
