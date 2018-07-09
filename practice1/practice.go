  判断两个map是否相等
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var s5 map[string]int
	s1 := make([]string, 3)
	s2 := make([]string, 3)
	s3 := make(map[string]int)
	s4 := make(map[string]int)
	for i := 0; i < 3; i++ {
		s1[i] = "xx"
		s2[i] = "x"
	}
	s3["cxx"] = 1
	s4["cxx"] = 1
	//s5["Golang"] = 1  零值map不能设置元素，在设置元素之前必须初始化map
	fmt.Println(reflect.DeepEqual(s1, s2))
	fmt.Println(reflect.DeepEqual(s3, s4))
	fmt.Println(len(s1))
}
/*package main

import "fmt"
*/
/*var ch chan int

func ts() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}

	*//*ch <- 0*//*
	<-ch
}

func main() {
	ch = make(chan int)
	go ts()
	ch <- 1
	for i := 0; i < 8; i++ {
		fmt.Println("er" + "i")
	}

	go ts()
	ch <- 1
	//time.Sleep(1 * time.Millisecond)
	*//*<-ch
	<-ch*//*

}*/
/*type A struct {
	a int
}
func main(){
	var a A
	h:=0
	a=A{a:2}
	b:=A{a:3}
	h,m:=2,3
	fmt.Println(b.a,a.a,h,m)
}*/
