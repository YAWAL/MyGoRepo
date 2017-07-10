package main

import "fmt"

func main()  {

	//Единственное отличие объявления среза от объявления массива — отсутствие указания длины в квадратных скобках.
	var x []float64
	fmt.Println(x)

	stringSlice := []string {"1", "2", "3", "4"}

	//fmt.Print("\n")
	fmt.Print(stringSlice)
	fmt.Print("\n")

	slice1 := make([]float64, 5)
	fmt.Println(slice1[3])

	//Для удобства мы также можем опустить low, high или и то, и другое. arr[0:] это то же самое что arr[0:len(arr)],
	// arr[:5] то же самое что arr[0:5] и arr[:] то же самое что arr[0:len(arr)]
	arr := [5]float64{1,2,3,4,5}
	slice2 := arr[2:4]
	fmt.Println(slice2)



}
