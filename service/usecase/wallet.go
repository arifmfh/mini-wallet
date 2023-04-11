package usecase

import (
	"github.com/arifmfh/go-mini-wallet/service"
)

type walletUsecase struct {
	WalletRepo service.WalletRepository
}

func WalletUsecase(walletRepo service.WalletRepository) service.WalletUsecase {
	return &walletUsecase{WalletRepo: walletRepo}
}
