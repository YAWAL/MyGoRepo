package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func a() {
	i := 4
	defer fmt.Println(i)
	i++
	i++
	return
}

//оператор defer, который позволяет отложить вызов указанной функции до тех пор, пока не завершится текущая
func main() {
	defer second()
	first()
	a()
}
//defer часто используется в случаях, когда нужно освободить ресурсы после завершения.
