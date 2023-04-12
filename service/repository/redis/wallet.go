package redis

import (
	"encoding/json"
	"time"

	"github.com/arifmfh/go-mini-wallet/models"
	"github.com/arifmfh/go-mini-wallet/service"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type walletRepository struct {
	RedisClient *redis.Client
}

func WalletRepository(redisClient *redis.Client) service.WalletRepository {
	return &walletRepository{RedisClient: redisClient}
}

func (w *walletRepository) Set(key string, val interface{}) (err error) {
	err = w.RedisClient.Set(key, val, time.Hour*24).Err()
	if err != nil {
		return err
	}

	return
}

func (w *walletRepository) Get(key string) (data string) {
	data, _ = w.RedisClient.Get(key).Result()
	return data
}

func (w *walletRepository) Register(costumerXID string) (err error) {
	id := uuid.New()
	wallet := models.Wallet{
		ID:      id.String(),
		OwnedBy: costumerXID,
		Status:  "disabled",
	}

	bytWallet, err := json.Marshal(wallet)
	if err != nil {
		return err
	}

	err = w.Set("wallet:"+costumerXID, string(bytWallet))
	if err != nil {
		return err
	}

	return
}

func (w *walletRepository) GetWallet(costumerXID string) (data models.Wallet, err error) {
	dataStr := w.Get(("wallet:" + costumerXID))
	err = json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (w *walletRepository) EnableWallet(param models.Wallet) (data models.Wallet, err error) {
	wallet := models.Wallet{
		ID:        param.ID,
		OwnedBy:   param.OwnedBy,
		Status:    "enabled",
		EnabledAt: time.Now().String(),
	}

	bytWallet, err := json.Marshal(wallet)
	if err != nil {
		return wallet, err
	}

	err = w.Set("wallet:"+param.OwnedBy, string(bytWallet))
	if err != nil {
		return wallet, err
	}

	return wallet, err
}
