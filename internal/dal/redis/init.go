package redis

import (
	"Dousheng_Backend/internal/dal/load"
	"Dousheng_Backend/utils/config"
	"Dousheng_Backend/utils/zap"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var DBRedis *redis.Client
var logger = zap.InitLogger(config.InitConfig("./config/logger/dal.yml"))

func init() {
	r := load.InitRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	// 测试链接ping-pong
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("failed to open db:%w\n", err))
	}
	fmt.Println(pong)

	DBRedis = rdb
	log.Println("Redis Connected!")
}
