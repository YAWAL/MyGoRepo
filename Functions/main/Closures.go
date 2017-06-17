package main

import "fmt"

//Возможно создавать функции внутри функций:

func main() {

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1,1))

	/*
	Функцию, использующую переменные, определенные вне этой функции, называют замыканием.
	 В нашем случае функция increment и переменная x образуют замыкание.
	 */
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
