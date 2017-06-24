package main

import (
	"sort"
	"fmt"
)

//binary search
func binary_search(array []int, item int) (string , int){
	low := 0
	high := len(array) - 1
 var k int
	for  k = 1; k <= 10; k++ {
		mid := low + high
		guess := array[mid]

		if guess == item {
			return "yes" , mid
		}
		if guess < item{
			high = mid - 1
		}else {
			low = mid + 1
		}
		return "absent", 0
	}



	return "none" , 0
}

func main()  {

	array := []int{1, 2 ,3 ,4 ,5, 6, 7, 8, 9}

	//fmt.Println(binary_search(array, 3))

	fmt.Println(sort.SearchInts(array, 456))
}