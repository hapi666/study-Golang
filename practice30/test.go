package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
}

var emp1 = Employee{1, "a", "aa"}

func EmplyeeByID(id int) Employee {
	return emp1
}

func main() {
	//EmplyeeByID(1).Name = "world" //编译报错
	a := EmplyeeByID(1)//可以正常运行
	fmt.Printf("%p",EmplyeeByID(1))
	fmt.Printf("\n%p",&a)
	fmt.Printf("%p",&emp1)

	a.Name = "hello"
}