package redis

import (
	"github.com/arifmfh/go-mini-wallet/service"
	"github.com/go-redis/redis"
)

type walletRepository struct {
	RedisClient *redis.Client
}

func WalletRepository(redisClient *redis.Client) service.WalletRepository {
	return &walletRepository{RedisClient: redisClient}
}
