package main

import (
    "fmt"
)

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    data1 := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data1 {
        defer v.print()
    }

    data2 := []*field{{"four"}, {"five"}, {"six"}}
    for _, v := range data2 {
        defer v.print()
    }
}