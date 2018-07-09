// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 1
// 	b := 2
// 	fmt.Println(a, b)
// 	a, b = b, a
// 	fmt.Println(a, b)

// }
package main

import (
     "fmt"
    "time"
)

func main() {
    var i int = 1

    go fmt.Println("result =>",func() int { return i * 2 }())
    i++
    time.Sleep(3*time.Second)
}
