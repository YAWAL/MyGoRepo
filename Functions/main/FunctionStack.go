package main

import "fmt"

func main()  {
	fmt.Println("main work...")

	f1()
}

func f1() int{
	fmt.Println("f1 work...")

	return f2()

}

func f2() int{
	fmt.Println("f2 work...")
	return 1
}
