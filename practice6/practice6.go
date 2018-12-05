package main

import (
	"fmt"
)

/*type csh int

func init() {
	csh = 6
}

func printcs() {
	fmt.Println(csh)
}*/

/*func main() {
	V := reflect.ValueOf("3")
	T := reflect.TypeOf("3")
	fmt.Println(V)
	fmt.Println(T)
}*/

/*type User struct {
}
type Admin struct {
	User
}

func (*User) ToString() {}

func (Admin) Test() {}
func main() {
	var u Admin
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}
	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u))
	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u))
}*/
/*type A struct {
	B B
	C C
}

type B struct {
	State int
}

type C struct {
	State int
}

func main() {
	a := make([]A, 5)
	var c C
	a[0].B.State, c.State = 1, 1
	var test = [5]int{5}
	fmt.Println(a[0].B.State)
	fmt.Println(c.State)
	fmt.Println(a)
	fmt.Println(test)
}*/

type mm interface {
	do()
}
type nn interface {
	notdo()
}

type s struct {
	s string
}

type t struct {
	s string
}

func (m *s) do() {
	m.s = m.s + "world!"
}

func (n *t) notdo() {
	n.s = n.s + " hapi!"
}

func (n *t) do() {
	n.s = n.s + " world!"
}

func main() {
	var str mm = &s{"hello "}
	var ttr nn = &t{"hi "}
	var super mm = &t{"hi "}
	if v, ok := str.(mm); ok {
		fmt.Println(v)
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
	fmt.Printf("%T\n%T\n%T", str, ttr, super)
}
