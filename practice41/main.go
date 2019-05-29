package main

import (
	"fmt"
)

func main(){
	array1 := [3][]string{
		[]string{"1","2","3"},
		[]string{"4","5","6"},
		[]string{"7","8","9"},
	}
	array2 := [3]string{"c","x","x"}
	array3 := []string{"h","a"}
	modifyArray1(array1)
	modifyArray2(array2)
	modifyArray3(array3)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}

func modifyArray1(array [3][]string){
	array[0][0] = "a"
	array[1] = []string{"4","5","a"}
}

func modifyArray2(array [3]string){
	array[0] = "h"
}

func modifyArray3(array []string){
	array[0] = "p"
}