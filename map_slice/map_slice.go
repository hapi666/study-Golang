package main

import "fmt"

func addList(s *[]int) {
	fmt.Printf("%p \n", s)
	q := make([]int, 0, 10)
	q = append(*s, 1)
	fmt.Printf("%p \n", q)
}

// func addMap(m map){

// }

func main() {
	t := make([]int, 0, 0)
	fmt.Printf("%p \n", t)
	addList(&t)
}
