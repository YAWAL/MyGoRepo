package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

var xs = []float64{12, 34, 45, 56}

func main()  {

	//xs := []float64{12, 34, 45, 56}

	fmt.Println(average(xs))
	
}
