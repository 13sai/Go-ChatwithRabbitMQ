package model

import (
	"fmt"
	"time"
	"bytes"
	"context"
	
	"github.com/spf13/viper"
	redis "github.com/go-redis/redis/v8"
)

const (
	Prefix = "im:"
)

var (
	RedisClient *redis.Client
	Ctx = context.Background()
)

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.MaxActive"),DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
	})

	pong, err := RedisClient.Ping(Ctx).Result()
	fmt.Println(pong, err)
}


// 获取key
func GetKey(key ...string) string {
	var bt bytes.Buffer
	bt.WriteString(Prefix)
	for _, arg := range key {
		bt.WriteString(arg)
    }
	//获得拼接后的字符串
	return bt.String()
}

// 记录日志pv，uv
func Logger(key string, userId string) {
	today := time.Now().Format("2006-01-02");

	pipe := RedisClient.Pipeline()
	fmt.Println(GetKey(key, today))
    // pipe.HIncrBy(Ctx, GetKey(key, today), today, 1)
    pipe.PFAdd(Ctx, GetKey("uv:", key, today), userId)
	
	// 总数
    pipe.HIncrBy(Ctx, GetKey("pv:", key), today, 1)
    _, err := pipe.Exec(Ctx)
    if err != nil {
        fmt.Println("err", err)
    }
}

func GetRedis() *redis.Client {
    return RedisClient
}