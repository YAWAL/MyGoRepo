package main

import "fmt"

func sendData2(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData2(chnl)
	fmt.Println(<-chnl)
}
