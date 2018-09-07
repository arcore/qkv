package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

//redis> SETEX cache_user_id 60 10086
//OK

//redis> GET cache_user_id  # 值
//"10086"

//redis> TTL cache_user_id  # 剩余生存时间
//(integer) 49
func main() {
	redisPool = newPool("192.168.16.200:8379", "1474741")
	conn := getRedis()
	defer conn.Close()
	log.Println(conn.Do("ZADD", "zset_test2", 1, "a"))
	log.Println(conn.Do("ZADD", "zset_test2", 2, "b"))
	log.Println(conn.Do("ZINCRBY", "zset_test2", 2, "b"))
	log.Println(conn.Do("ZADD", "zset_test2", 456456, "c"))
	log.Println(redis.Int(conn.Do("EXPIRE", "zset_test2", 10)))
	log.Println(redis.Strings(conn.Do("ZRANGE", "zset_test2", 0, -1, "WITHSCORES")))
	time.Sleep(time.Second * 11)
	log.Println(redis.Strings(conn.Do("ZRANGE", "zset_test2", 0, -1, "WITHSCORES")))

}
func getRedis() redis.Conn {
	return redisPool.Get()
}
func newPool(host, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 256, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					return nil, err
				}
			}
			return c, err
		},
	}
}
