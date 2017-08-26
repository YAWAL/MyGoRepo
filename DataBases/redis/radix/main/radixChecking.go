package main

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"fmt"
	"os"
)



func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}



func main() {
	conn, err := GetRedisConn(NETWORK, ADDRESS, POOL_SIZE)
	if err != nil{
		loger.Log.Error(err)
	}

	// select database
	r := conn.Cmd("select", 0)
	errHndlr(r.Err)

	r = conn.Cmd("echo", "Hello world!")
	errHndlr(r.Err)

	r = conn.Cmd("set", "mykey0", "myval0")
	errHndlr(r.Err)

	s, err := conn.Cmd("get", "mykey0").Str()
	errHndlr(err)
	fmt.Println("mykey0:", s)


}
