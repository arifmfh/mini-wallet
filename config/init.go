package config

import (
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

type Config struct {
	Router      *chi.Mux
	RedisClient *redis.Client
}

func Init() *Config {
	var cfg Config

	err := cfg.initRedis()
	if err != nil {
		panic(err.Error())
	}

	err = cfg.initChi()
	if err != nil {
		panic(err.Error())
	}

	err = cfg.initService()
	if err != nil {
		panic(err.Error())
	}

	return &cfg
}
