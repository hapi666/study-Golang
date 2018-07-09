package main

import "fmt"

type NewString string

type getter interface{
	Get() interface{}
}

func (tt NewString) Get() interface{} {
	return string(tt)
}

func testinterface(gg getter)  {
	fmt.Println(gg.Get())
}

func main()  {
	ch:=make(chan string)
	newstring:=NewString("haha")
	testinterface(newstring)
}

