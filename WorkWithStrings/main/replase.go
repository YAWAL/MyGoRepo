package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 0))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 1))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 2))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", -1))//If n is less than 0, there is no limit on the number of replacements.
	fmt.Println(strings.Replace("1234534232132", "1", "0", -1))//If n is less than 0, there is no limit on the number of replacements.
	// case-sensitive
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "AN", "anese", -1))
}
