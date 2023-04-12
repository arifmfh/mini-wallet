package service

import "github.com/arifmfh/go-mini-wallet/models"

type WalletUsecase interface {
	Register(costumerXID string) (err error)
	EnableWallet(costumerXID string) (data models.Wallet, code int, err error)
	GetWallet(costumerXID string) (data models.Wallet, code int, err error)
}
