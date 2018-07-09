package main

import (
	"fmt"
)

//可见slice的切割是左闭右开的
func main() {
	s := make([]int, 10)
	s[0] = 2
	s[1] = 3
	s[2] = 4
	s[3] = 5
	ss := s[0:3]
	sss := s[1:3]
	fmt.Println(ss)
	fmt.Println(sss)
	fmt.Println(s)
}
