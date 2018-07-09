package main

import (
	"fmt"
	"strconv"
)

func main() {
	ss := "666"
	in, err := strconv.Atoi(ss) //strconv.Atoi是把只包含数字的字符串转换成int型数字
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(in)
	}
	s := "hapi666"
	c := []byte(s)
	c[0] = 'c'
	s = string(c)
	fmt.Println(s)
}
