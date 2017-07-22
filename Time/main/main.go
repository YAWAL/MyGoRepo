package main

import (
	"strconv"
	"time"
	"fmt"
)

var timestamp = strconv.FormatInt(time.Now().Unix(), 10)

func showTimestamp()  {
	fmt.Printf(timestamp)

}
func main()  {
	showTimestamp()
}