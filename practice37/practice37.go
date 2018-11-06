package main

import "fmt"

func main(){
	a:=make([]int,10)
	b:=len(a)>>1
	fmt.Println(b)
}
