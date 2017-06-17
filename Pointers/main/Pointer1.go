package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}
/*
Указатели указывают на участок в памяти, где хранится значение.
 Используя указатель (*int) в функции zero, мы можем изменить значение оригинальной переменной.
В Go указатели представлены через оператор * (звёздочка), за которым следует тип хранимого значения.
 В функции zero xPtr является указателем на int.

 оператор &, который используется для получения адреса переменной. &x вернет *int (указатель на int)
  потому что x имеет тип int.
 */