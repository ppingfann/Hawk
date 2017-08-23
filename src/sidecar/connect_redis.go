package sidecar

import "github.com/garyburd/redigo/redis"

//connect to redis
func Conn_redis() redis.Conn{
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	return conn
}

	var Conn_redi redis.Conn = Conn_redis()
