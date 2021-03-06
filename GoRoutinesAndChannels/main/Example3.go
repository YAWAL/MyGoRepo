package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done

	go hello(done)

	fmt.Println("main function")
}
