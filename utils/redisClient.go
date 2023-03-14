package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInit(Host string, Port int, Password string, DB int) (*redis.Client, error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Host, Port),
		Password: Password,
		DB:       DB,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}

	return RedisClient, err
}
