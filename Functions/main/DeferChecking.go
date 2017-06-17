package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

//оператор defer, который позволяет отложить вызов указанной функции до тех пор, пока не завершится текущая
func main() {
	defer second()
	first()
}
//defer часто используется в случаях, когда нужно освободить ресурсы после завершения.
