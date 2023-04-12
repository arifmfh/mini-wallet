package usecase

import (
	"fmt"

	"github.com/arifmfh/go-mini-wallet/models"
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
		return err
	}

	return
}

func (w *walletUsecase) EnableWallet(costumerXID string) (data models.Wallet, code int, err error) {
	wallet, err := w.WalletRepo.GetWallet(costumerXID)
	if err != nil {
		return data, 500, err
	}

	if wallet.Status == "disabled" {
		data, err = w.WalletRepo.EnableWallet(wallet)
		if err != nil {
			return data, 500, err
		}

		return data, 201, nil

	} else if wallet.Status == "enabled" {
		return data, 400, fmt.Errorf("Already enabled")
	} else {
		return data, 400, fmt.Errorf("Wallet not found")
	}
}

func (w *walletUsecase) GetWallet(costumerXID string) (data models.Wallet, code int, err error) {
	wallet, err := w.WalletRepo.GetWallet(costumerXID)
	if err != nil {
		return data, 500, err
	}

	if wallet.ID == "" {
		return data, 400, fmt.Errorf("Wallet not found")
	}

	return wallet, 200, nil
}

func (w *walletUsecase) Deposit(param models.Deposit) (data models.Deposit, code int, err error) {
	wallet, err := w.WalletRepo.GetWallet(param.DepositedBy)
	if err != nil {
		return data, 500, err
	}

	if wallet.ID == "" {
		return data, 400, fmt.Errorf("Wallet not found")
	}
	if wallet.Status != "enabled" {
		return data, 400, fmt.Errorf("Wallet is disabled")
	}

	wallet.Balance += param.Amount

	data, err = w.WalletRepo.Deposit(wallet, param)
	if err != nil {
		return data, 500, err
	}

	return data, 200, nil
}
