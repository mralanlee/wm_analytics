package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

type Subscriber struct {
	pubsub   *redis.PubSub
	callback handler
}

type handler func(string)

var (
	ctx        = context.Background()
	redis_port = os.Getenv("REDIS_PORT")
	redis_host = os.Getenv("REDIS_ADDRESS")
)

var Redis *RedisClient

func init() {
	redisAddress := fmt.Sprintf("%s:%s", redis_host, redis_port)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	Redis = &RedisClient{client}
}

// func NewSub(channel string) (*Subscriber, error) {
// 	s := Subscriber{
// 		pubsub: Redis.client.Subscribe(ctx, channel),
// 	}
// }
//
