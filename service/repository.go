package service

type WalletRepository interface {
	Register(costumerXID string) (err error)
}
