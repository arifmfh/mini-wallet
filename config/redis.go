package config

import (
	"os"

	"github.com/go-redis/redis"
)

func (c *Config) initRedis() error {
	c.RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	return nil
}
