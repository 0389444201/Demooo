package rediscache

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
