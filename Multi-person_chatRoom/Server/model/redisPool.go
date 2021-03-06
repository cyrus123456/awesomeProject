package model

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitPool(maxIdle, maxActive int, idleTimeout time.Duration, DialAdress string) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲链接数
		MaxActive:   maxActive,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接的代码，链 接哪个ip的redis
			return redis.Dial("tcp", DialAdress)
		},
	}

}
