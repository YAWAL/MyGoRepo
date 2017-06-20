package main

import (
	"hash/crc32"
	"fmt"
)

//Обычным применением crc32 является сравнение двух файлов
func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}
