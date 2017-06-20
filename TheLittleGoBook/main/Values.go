package main

import "fmt"


func log(message string) {
}


func add(a int, b int) int {
	return a + b
}
	func power(name string) (int, bool) {

		return 5, true
}


func main()  {


	value, exists := power("goku")

	fmt.Println(value, exists)
}

