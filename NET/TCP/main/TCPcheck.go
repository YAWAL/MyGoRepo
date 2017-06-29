package main

import (
	"net"
	"fmt"
	"bufio"
)

func main(){

	listener, _ := net.Listen("tcp", ":5000")

	for {
		conn, err := listener.Accept()

		if err != nil{
			fmt.Println("Can't connect")
			conn.Close()
			continue
		}

		fmt.Println("connected")

		bufReader := bufio.NewReader(conn)

		go func(conn net.Conn) {

			defer conn.Close()

			for {
				rbyte, err := bufReader.ReadByte()

				if err != nil {
					fmt.Println("Cant read", err)
					break
				}
				fmt.Print(string(rbyte))
			}
		}(conn)
	}


}
