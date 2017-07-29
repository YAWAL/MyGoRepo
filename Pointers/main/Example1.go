package main

import (
	"fmt"
)


func main() {
	b := 255
	var a *int = &b
	var c  = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Type of c is %T\n", c)
	fmt.Println("address of b is", a)
	fmt.Println("address of c is", c)

}