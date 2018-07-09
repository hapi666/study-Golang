package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str := "hapi666"
	s := "p"
	fmt.Println(strings.Index(str, s))
	fmt.Println(strings.Replace(str, s, "cxx", 4))
	//fmt.Println(time.Time.UTC)
}
