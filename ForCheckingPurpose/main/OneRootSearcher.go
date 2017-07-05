package main

import "fmt"

//answer for task 243а page 104
func SearchOneRoot(n int) bool{
	var x, y int
	rootCounter := 0
	for x = 1; x < n; x ++ {
		for y = 1; y < n ; y ++  {
			if x*x + y*y == n && y <= x{//
				rootCounter++
				fmt.Println("Task a -> Pair of natural numbers for n =" , n, "when x^2 + y^2 = n : x =", x, " y =", y)
			}
			if rootCounter > 1 {
				break
			}
		}
		return true
	}
	if rootCounter == 0{
		fmt.Println("there is no root for n =", n)
	}
	return false
}

