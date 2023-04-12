package service

import "github.com/arifmfh/go-mini-wallet/models"

type WalletUsecase interface {
	Register(costumerXID string) (err error)
	EnableWallet(costumerXID string) (data models.Wallet, code int, err error)
	GetWallet(costumerXID string) (data models.Wallet, code int, err error)
	GetTransactions(costumerXID string) (data []models.Transaction, code int, err error)
	Deposit(param models.Deposit) (data models.Deposit, code int, err error)
	Withdraw(param models.Withdraw) (data models.Withdraw, code int, err error)
}
