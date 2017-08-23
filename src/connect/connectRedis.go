package connect

import "github.com/garyburd/redigo/redis"

//connect to redis
func ConnRedis() redis.Conn{
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	return conn
}

	var ConnRedi redis.Conn = ConnRedis()
