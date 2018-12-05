package main

import (
	"fmt"
)

func main() {
	var s = [5]string{"1", "2", "3", "4", "5"}
	fmt.Printf("原数组s:\n %v\n", s)

	sli := s[0:5]
	newsli := append(sli, "t")
	sli[1] = "x"
	fmt.Printf("截取数组s后的sli:\n %v\n", sli)
	fmt.Printf("原数组s:\n %v\n", s)
	fmt.Printf("append sli后的newsli:\n %v\n", newsli)

	newsli[0] = "w"
	fmt.Printf("改变了newsli的第一个元素之后的newsli:\n %v\n", newsli)
	fmt.Printf("原数组:\n %v\n", s)
	fmt.Printf("sli:\n %v\n", sli)
}
