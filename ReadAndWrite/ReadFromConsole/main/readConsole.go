package main

import "fmt"

func main() {
	var i int
	fmt.Println("enter number")

	i, err := fmt.Scanf( &i, "%d")

	fmt.Println("i = ", i)

	if err != nil {
		fmt.Println("i = ", i)
	}
	fmt.Println(err)

}