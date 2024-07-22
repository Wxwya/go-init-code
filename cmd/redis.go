package cmd

import (
	"fmt"
	"xwya/server"
	"xwya/utils"

	"github.com/gomodule/redigo/redis"
)

func InitRedis() {
	server.Pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 表示和数据库的最大连接数, 0 表示没有限制
		IdleTimeout: 100, // 表示一个连接在空闲状态超过100秒 ,就会关闭, 默认是
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", utils.RedisHostAndPort, redis.DialPassword(utils.RedisPassword))
			if err != nil {
				return nil, fmt.Errorf("failed to connect to Redis: %w", err)
			}
			return conn, nil
		},
	}
}

func GetConn(db int) *redis.Conn {
	conn := server.Pool.Get()
	if _, err := conn.Do("SELECT", db); err != nil {
		WriteLog(err)
		return nil
	}
	return &conn
}
func CloseRedis() {
	if server.Pool != nil {
		if err := server.Pool.Close(); err != nil {
			WriteLog(err)
			return
		}
		fmt.Println("redis连接已关闭")
	}
}
