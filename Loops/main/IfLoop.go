package main

import "fmt"

func IfLoop()  {

	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i , " odd")
		} else{
			fmt.Println(i , " even")
		}

	}

	return
}
