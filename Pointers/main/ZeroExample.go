package main

import "fmt"

func zero(x int) {
	x = 0
}
func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x всё еще равен 5
}
/*
В этой программе функция zero не изменяет оригинальную переменную x из функции main.
 */