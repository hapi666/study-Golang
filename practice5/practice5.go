package main

import (
	"fmt"
)

func func_name(str string) string {
	return str
}

func main() {
	/*var str string
	str = "23"*/
	/*str2 := []string{"2", "3"}
	str2[0] = "56"
	fmt.Println(str[0]) //2的ascii码
	fmt.Println(str)
	fmt.Println(func_name(str))
	fmt.Printf("%p", str2)*/
	/*fmt.Println(test.csh)*/
	if "123" == "123" {
		fmt.Println("yes")
	}
	s := make(int, 10)
	s[0] = 2
	s[1] = 3
	s[2] = 4
	s[3] = 5
	ss := s[:3]
	fmt.Println(ss)
	fmt.Println(s)
}
