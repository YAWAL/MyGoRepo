package main

import (
	"math"
	"fmt"

)

/*
Дано ціле число m > 1. Отримати найбільше ціле k, при якому 4^k < m
page 44
 */

func ShowBiggestK( m int) {
	if  IsMmoreThan5(m){
		fmt.Println("Biggest k value for equation 4^k < m , where m = ", m, " -> ", findMaxValueOfK(CreateSliceOfK(m)))
	}
}

func IsMmoreThan5(m int) bool {
	if m > 5 {
		return true
	}
	fmt.Println("There are no k where equation 4^k < m is true")
	return false
}

func CreateSliceOfK(m int) []int{
	//empty slice for allocation values of k with capacity m (to prevent runtime error -> panic: runtime error: index out of range)
	valuesOfK := make([]int, m)
	//creating slice of k
	for k := 1; int(math.Pow(4, float64(k))) < m; k++{
		for j := k - 1; j < k ; j++ {
			valuesOfK[j] = k
		}
	}
	return valuesOfK
}

func findMaxValueOfK(valuesOfK []int) int {
	max := valuesOfK[0]
	for _, next := range valuesOfK {
		if max < next  {
			max = next
		}
	}
	return max
}

func main()  {
	fmt.Println("_____________ \n Дано ціле число m > 1. Отримати найбільше ціле k, при якому 4^k < m")
	for   {
		var m int
		fmt.Println("Please, enter m : ")
		fmt.Scanln(&m)
		ShowBiggestK(m)
	}
}