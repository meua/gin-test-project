package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, e := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if e != nil {
		fmt.Errorf("connect to redis failed, err is:", e)
	}
	reply, err := conn.Do("lpush", "mobile:p30:n3000:1", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if err != nil {
		fmt.Errorf("operate to redis failed, err is:", e)
	}
	println(reply)

}
