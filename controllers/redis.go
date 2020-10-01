package pubsub

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/mralanlee/wm_analytics/common"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisAddr     string
	RedisPassword string
}

type RedisClient struct {
	Client *redis.Client
}

type Subscriber struct {
	pubsub   *redis.PubSub
	callback handler
}

type handler func(*redis.Message)

var (
	ctx        = context.Background()
	redis_port = os.Getenv("REDIS_PORT")
	redis_host = os.Getenv("REDIS_ADDRESS")
	redis_pass = os.Getenv("REDIS_PASSWORD")
	Redis      *RedisClient
)

func (c *RedisConfig) fill() {
	if c.RedisHost == "" {
		c.RedisHost = common.REDIS_HOST
	}

	if c.RedisPort == "" {
		c.RedisPort = common.REDIS_PORT
	}

	if c.RedisPassword == "" {
		c.RedisPassword = common.REDIS_PASSWORD
	}

	c.RedisAddr = fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func (r *RedisClient) NewSubscriber(channel string, cb handler) *Subscriber {
	sub := Subscriber{
		pubsub:   r.Client.Subscribe(ctx, channel),
		callback: cb,
	}

	return &sub
}

func (s *Subscriber) Listen() {
	go func() {
		for {
			msg, _ := s.pubsub.ReceiveMessage(ctx)

			go func() {
				s.callback(msg)
			}()
		}
	}()
}

func init() {
	var config = &RedisConfig{
		RedisHost:     redis_host,
		RedisPort:     redis_port,
		RedisPassword: redis_pass,
	}

	config.fill()

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	fmt.Println("Redis initialized")
	Redis = &RedisClient{Client: client}
}
