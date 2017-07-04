package main

import (
	"fmt"
	"strings"
)
func checkContains(){
	row1 := "asdfghjklzx{}}{{Pcvbm"
	row2 := "{"
	if strings.Contains(row1, row2) {
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}

}
func main() {
	checkContains()
}
