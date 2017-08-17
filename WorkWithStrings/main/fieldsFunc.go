package main

import (
	"strings"
	"fmt"
	"unicode"
)

func main() {

	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	strArray := strings.FieldsFunc(`Australia major cities – Sydney, Brisbane,
                                 Melbourne, Perth, Adelaide – are coastal`,x)
	fmt.Println(strArray)
	for _, v := range strArray {
		fmt.Println(v)
	}

	fmt.Println("\n*****************Split by number*******************\n")

	y := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	testArray := strings.FieldsFunc(`1 Sydney Opera House.2 Great Barrier Reef.3 Uluru-Kata Tjuta National Park.4 Sydney Harbour Bridge.5 Blue Mountains National Park.6 Melbourne.7 Bondi Beach`,y)
	for _, w := range testArray {
		fmt.Println(w)
	}

}
