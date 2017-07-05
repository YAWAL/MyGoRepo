package main

import (
	"fmt"
)
func main()  {

	WriteTask()

	for   {
		var n int
		fmt.Println("Please, enter n : ")
		fmt.Scanln(&n)
		SearchOneRoot(n)
		SearchAllRoots(n)
	}
}
