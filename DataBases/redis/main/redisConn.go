package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func GetRedisConnection() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}
func main() {
	connection,_ := dial()
	defer connection.Close()
	data, err := redis.Values(connection.Do("SORT", "Users", "BY", "User:*->name","GET", "User:*->name"))
	if (err) {
		fmt.Println("Error getting values", err)
	}
	for i:= range data {
		var Uname string
		data,err := redis.Scan(data, &Uname)
		if (err) {
			fmt.Println("Error getting value",err)
		}else {
			fmt.Println("Name Uname")
		}
	}
}