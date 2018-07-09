package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(zs int)
}

type Simple struct {
	value int
}

func (s *Simple) Get() int {
	return s.value
}

func (ss *Simple) Set(zs int) {
	ss.value = zs
}

func JK(jk Simpler) {
	q := jk.Get()
	fmt.Println(q)
	jk.Set(6)
	fmt.Println(jk.Get())
}

func main() {
	var testinterface Simpler = &Simple{7}
	JK(testinterface)
	switch t := testinterface.(type) {
	case *Simple:
		fmt.Printf("Type Simple %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
