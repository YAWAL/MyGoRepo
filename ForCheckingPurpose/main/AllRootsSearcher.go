package main

import "fmt"

//answer for task 243б page 104
func searchAllRoots(n int)  {
	var x, y int
	rootCounter := 0
	for x = 1; x < n; x ++ {
		for y = 1; y < n ; y ++  {
			if x*x + y*y == n && y <= x {
				rootCounter++
				fmt.Println("Task б -> Pairs of natural numbers for n =" , n, "when x^2 + y^2 = n and x >= y: x =", x, " y =", y)
			}
		}
	}
}
