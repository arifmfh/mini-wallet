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

func (w *walletUsecase) GetTransactions(costumerXID string) (data []models.Transaction, code int, err error) {
	wallet, err := w.WalletRepo.GetWallet(costumerXID)
	if err != nil {
		return data, 500, err
	}

	if wallet.ID == "" {
		return data, 400, fmt.Errorf("Wallet not found")
	}
	if wallet.Status != "enabled" {
		return data, 400, fmt.Errorf("Wallet is disabled")
	}

	data, err = w.WalletRepo.GetTransactions(costumerXID)
	if err != nil {
		return data, 500, err
	}

	if len(data) == 0 {
		return data, 400, fmt.Errorf("There is no transaction")
	}

	return data, 200, nil
}

func (w *walletUsecase) Deposit(param models.Deposit) (data models.Deposit, code int, err error) {
	isDuplicate, err := w.WalletRepo.DepositCheckReferenceID(param.ReferenceID)
	if err != nil {
		return data, 500, err
	}
	if isDuplicate {
		return data, 400, fmt.Errorf("Duplicate reference_id")
	}

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

func (w *walletUsecase) Withdraw(param models.Withdraw) (data models.Withdraw, code int, err error) {
	isDuplicate, err := w.WalletRepo.WithdrawCheckReferenceID(param.ReferenceID)
	if err != nil {
		return data, 500, err
	}
	if isDuplicate {
		return data, 400, fmt.Errorf("Duplicate reference_id")
	}

	wallet, err := w.WalletRepo.GetWallet(param.WithdrawnBy)
	if err != nil {
		return data, 500, err
	}

	if wallet.ID == "" {
		return data, 400, fmt.Errorf("Wallet not found")
	}
	if wallet.Status != "enabled" {
		return data, 400, fmt.Errorf("Wallet is disabled")
	}

	wallet.Balance -= param.Amount
	if wallet.Balance < 0 {
		return data, 400, fmt.Errorf("Balance not sufficient")
	}

	data, err = w.WalletRepo.Withdraw(wallet, param)
	if err != nil {
		return data, 500, err
	}

	return data, 200, nil
}
