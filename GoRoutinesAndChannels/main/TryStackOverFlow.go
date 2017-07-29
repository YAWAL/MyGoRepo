package main

import (
	"fmt"
)


func createGorutine(){


	fmt.Println("thread # ")
}

func main()  {
	for i := 0; i <= 1000000000; i++ {
		go createGorutine()

		fmt.Println("thread # ", i)

	}
	fmt.Println("main function")
}
