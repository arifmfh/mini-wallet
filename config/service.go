package config

import (
	"github.com/arifmfh/go-mini-wallet/service/delivery/http"
	"github.com/arifmfh/go-mini-wallet/service/repository/redis"
	"github.com/arifmfh/go-mini-wallet/service/usecase"
)

func (cfg *Config) initService() error {

	walletRepo := redis.WalletRepository(cfg.RedisClient)

	walletUsecase := usecase.WalletUsecase(walletRepo)

	http.Router(cfg.Router, walletUsecase)

	return nil
}
