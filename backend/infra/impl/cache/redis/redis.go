package redis

import (
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/crazyfrankie/ddd-todolist/backend/conf"
)

type Client = redis.Client

func New() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Addr, // Redis地址
		DB:       0,                         // 默认数据库
		Password: conf.GetConf().Redis.Password,
		// connection pool configuration
		PoolSize:        100,             // Maximum number of connections (recommended to set to CPU cores * 10)
		MinIdleConns:    10,              // minimum idle connection
		MaxIdleConns:    30,              // maximum idle connection
		ConnMaxIdleTime: 5 * time.Minute, // Idle connection timeout

		// timeout configuration
		DialTimeout:  5 * time.Second, // Connection establishment timed out
		ReadTimeout:  3 * time.Second, // read operation timed out
		WriteTimeout: 3 * time.Second, // write operation timed out
	})

	return rdb
}
