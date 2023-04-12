package service

import "github.com/arifmfh/go-mini-wallet/models"

type WalletRepository interface {
	Register(costumerXID string) (err error)
	GetWallet(costumerXID string) (data models.Wallet, err error)
	EnableWallet(param models.Wallet) (data models.Wallet, err error)
	GetTransactions(costumerXID string) (data []models.Transaction, err error)
	DepositCheckReferenceID(referenceID string) (IsDuplicate bool, err error)
	Deposit(wallet models.Wallet, param models.Deposit) (data models.Deposit, err error)
	WithdrawCheckReferenceID(referenceID string) (IsDuplicate bool, err error)
	Withdraw(wallet models.Wallet, param models.Withdraw) (data models.Withdraw, err error)
}
