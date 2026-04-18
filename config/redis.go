package config

import (
    "context"
    "os"

    "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func InitRedis() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_ADDR"),     
        Password: os.Getenv("REDIS_PASSWORD"), 
        DB:       0,
    })
    return rdb
}