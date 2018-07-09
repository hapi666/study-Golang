package main

import (
	"fmt"
)

func test1(ch chan interface{})  {
	value:=<-ch
	fmt.Println(value.(string))
}

func test2(ch interface{})  {
	value:=<-ch.(chan string)
	fmt.Println(value)
}

func test23(ch interface{})  {
	value:=<-ch.(chan interface{})
	fmt.Println(value.(string))
}

func main()  {
	T1:=make(chan string,2)
	T2:=make(chan interface{},3)
	T1<-"hapi666"
	T1<-"666"
	T2<-"cxx666"
	T2<-"111"
	T2<-"222"
	test1(T2)
	test2(T1)
	test23(T2)
	//test3(T1)

}
