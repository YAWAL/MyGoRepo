package main

import "fmt"

func main()  {

	//var stringMap map[string]string
	//var intMAp map[int]int

	stringIntMap := make(map[string]int)
	stringIntMap["key"] = 1
	stringIntMap["key1"] = 2
	stringIntMap["key2"] = 3
	stringIntMap["key3"] = 4
	stringIntMap["key4"] = 5


	//fmt.Println(stringIntMap["key"])
	//fmt.Print(stringIntMap["key1"])
	fmt.Print(stringIntMap)




}
