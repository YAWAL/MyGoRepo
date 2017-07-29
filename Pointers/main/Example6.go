package main

import "fmt"

func modify1(arr *[3]int) {
	(*arr)[0] = 90
}

func modify2(arr *[3]int) {
	arr[0] = 90
}

func main() {
	a := [3]int{89, 90, 91}
	modify1(&a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
