package service

type WalletUsecase interface {
	Register(costumerXID string) (err error)
}
