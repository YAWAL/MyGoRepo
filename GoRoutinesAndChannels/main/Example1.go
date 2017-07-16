package main

import (
	"fmt"
	"time"
)

func foo(d time.Duration, c chan int) {
	d *= 1000000000
	time.Sleep(d)
	fmt.Println(d)
	c <- 1
}

func main() {
	c := make(chan int)
	go foo(3, c)
	go foo(2, c)
	go foo(1, c)
	<- c
	<- c
	<- c
}
