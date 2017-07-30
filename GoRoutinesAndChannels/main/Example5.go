package main

import (
	"fmt"
	"time"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	startTime := time.Now().Nanosecond()
	go calcSquares(number, sqrch)
	firsGoStart := time.Now().Nanosecond()
	go calcCubes(number, cubech)
	secondGoStart := time.Now().Nanosecond()

	squares, cubes := <-sqrch, <-cubech
	endTime := time.Now().Nanosecond()

	fmt.Println("Final output", squares + cubes, "time : ", endTime - startTime, " diff time : ", secondGoStart - firsGoStart)
}
