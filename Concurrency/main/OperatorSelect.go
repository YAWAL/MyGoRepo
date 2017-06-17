package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			//select {
			//case msg1 := <- c1:
			//	fmt.Println(msg1)
			//case msg2 := <- c2:
			//	fmt.Println(msg2)
			//}

			select {
			case msg1 := <- c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <- c2:
				fmt.Println("Message 2", msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}

		}
	}()

	var input string
	fmt.Scanln(&input)
}
