package main

import "fmt"

func main()  {

	var MyArray [5]int
	fmt.Println(MyArray)

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	//x := [5]float64{ 98, 93, 77, 82, 83 }

	y := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0

	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(y)
	fmt.Println(total / float64(len(y)))


	for _ , value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

}
