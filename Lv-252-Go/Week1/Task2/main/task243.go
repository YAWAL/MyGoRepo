package main

import "fmt"

//answer for task 243б page 104
func twoRootsFor243(n int)  {
	var x, y int
	for x = 1; x < n; x ++ {
		for y = 1; y < n; y ++  {
			if x*x + y*y == n {
				//fmt.Println("x*x + y*y = n" , x, y ,n)
				fmt.Println("twoRootsFor243:Pair of natural numbers when x^2 + y^2 = n : x =", x, " y =", y, " n = ", n)
			}
		}
	}
}

//answer for task 243а page 104
func oneRootFor243(n int)  {
	var x int
	//when one root (y) member already defined (y = 1)
	y := 1
	for x = 1; x < n; x ++ {
		if x*x + y*y == n {
			//fmt.Println("x*x + y*y = n" , x, y ,n)
			fmt.Println("oneRootFor243:Pair of natural numbers when x^2 + y^2 = n : x =", x, " y =", y, " n = ", n)
		}
	}
}
func main()  {
	oneRootFor243(10)
	twoRootsFor243(10)
}
