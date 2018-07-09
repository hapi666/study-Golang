package main

import(
	"fmt"
)

// func test(p *int) {
// 	fmt.Printf("%p\n",p)
// 	*p=6
// }

// func main()  {
// 	i:=10
// 	fmt.Printf("%p\n",&i)
// 	p := &i
// 	fmt.Printf("%p\n",&p)
// 	test(p)
// 	fmt.Println(i)
// 	fmt.Printf("%p",&i)
// }

// func test(p int) {
// 	p=9
// 	fmt.Printf("%p\n",&p)
// }

// func main()  {
// 	i:=10
// 	fmt.Printf("%p\n",&i)
// 	test(i)
// 	fmt.Println(i)
// }

func test(m map[string]int)  {
	m["hapi"] = 16
	fmt.Printf("%p\n",&m["hapi"])
}

func main()  {
	mm:=make(map[string]int)
	mm["hapi"] = 21
	fmt.Println(mm["hapi"])
	fmt.Printf("%p\n",&mm["hapi"])
	test(mm)
	fmt.Println(mm["hapi"])
}