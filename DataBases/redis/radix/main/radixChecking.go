package main

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"fmt"
	"os"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

const NETWORK = "tcp"
const ADDRESS = "localhost:6379"
const POOL_SIZE = 10

func GetRedisConn(network string, addr string, size int) (conn *redis.Client, err error) {
	p, err := pool.New(network, addr, size)
	if err != nil {
		loger.Log.Error(err)
	}
	loger.Log.Print("pool created")
	conn, err = p.Get()
	if err != nil {
		loger.Log.Error(err)
	}
	loger.Log.Print("connection established")
	//resp := conn.Cmd("ping")
	//fmt.Printf("resi value is %v , type is %T \n", resp, resp )
	return conn, err
}

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	conn, err := GetRedisConn(NETWORK, ADDRESS, POOL_SIZE)
	if err != nil {
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

	myhash := map[string]string{
		"mykey1": "myval1",
		"mykey2": "myval2",
		"mykey3": "myval3",
	}

	// Alternatively:
	// c.Cmd("mset", "mykey1", "myval1", "mykey2", "myval2", "mykey3", "myval3")
	r = conn.Cmd("mset", myhash)
	errHndlr(r.Err)

}
