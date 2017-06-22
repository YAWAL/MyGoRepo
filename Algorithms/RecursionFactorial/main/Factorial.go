package main

//import "fmt"

func factorial(i int) int {
	var fact int
	if i == 1 {
		fact = 1
	}else{
		fact = i * factorial(i - 1)
	}
	return fact
}

func printFactorial(i int)  {
	if i == 1 {
		println("factorial 1! = 1")
	}else {
		println("factorial " , i, "! = ", factorial(i))
	}

}

func main()  {
	i := 7
	factorial(i)
	printFactorial(i)
}
