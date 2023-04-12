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

func (w *walletUsecase) Register(costumerXID string) (err error) {
	err = w.WalletRepo.Register(costumerXID)
	if err != nil {
		return nil
	}

	return
}
