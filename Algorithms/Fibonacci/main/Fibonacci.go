package main

import "fmt"

func fibonacci(i int) int{

sum := 0
	if i == 0 || i ==1 {
		return 1
	}else{
		for k := i; k == i; i++ {
			sum = fibonacci(i - 1) + fibonacci(i - 2)
		}
		 return sum
	}
}



func main()  {

	fmt.Println("fibonacci(i - 1) + fibonacci(i - 2) = ", fibonacci(21) )
}
