package main

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"fmt"
)

func main() {
	p, err := pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		loger.Log.Error(err)
	}
	loger.Log.Print("pool created")

	// In another go-routine

	conn, err := p.Get()
	if err != nil {
		loger.Log.Error(err)
	}
	loger.Log.Print("connection established")


	resp := conn.Cmd("ping")
	fmt.Printf("resi value is %v , type is %T", resp, resp )

	if conn.Cmd("SADD", "myset", "some").Err != nil {
		loger.Log.Errorf("error : ", err)

	}

	if conn.Cmd("SOME", "CMD").Err != nil {
		loger.Log.Errorf("error : ", err)
	}

	p.Put(conn)

}
