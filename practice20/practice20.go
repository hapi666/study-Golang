package main

import (
	"bytes"
	"fmt"
)

//复用slice

//做个数组，然后再切个片来用，没超1000的时候就是复用，超了再分配。
//s = buffer[:0] 复用之前务必确认所有指向buffer的切片、切过的切片都断开引用，不再指向buffer
func main() {
	buffer := [1000]byte{}
	s := buffer[:0]
	fmt.Println(s)
	var buffer2 bytes.Buffer
	for i:=0; i<10;i++  {
		buffer2.WriteString("1")
	}
	fmt.Println(buffer2.String())
}
