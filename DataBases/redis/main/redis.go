package main

import "github.com/garyburd/redigo/redis"

var conn redis.Conn

func main()  {
	defer conn.Close()
	n, err := conn.Do("APPEND", "key", "value")
	if err != nil{

	}
	conn.Do()
	println(n)
}
