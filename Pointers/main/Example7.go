package main

func main() {
	b := [...]int{109, 110, 111}
	p := &b
	p++

	/*
	The above program will throw compilation error main.go:6: invalid operation: p++ (non-numeric type *[3]int)
	 */
}
