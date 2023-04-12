package http

import (
	"github.com/arifmfh/go-mini-wallet/service"
	"github.com/go-chi/chi"
)

type Handler struct {
	WalletUsecase service.WalletUsecase
}

func Router(r *chi.Mux, walletUsecase service.WalletUsecase) {
	h := &Handler{
		WalletUsecase: walletUsecase,
	}

	r.Route("/api/v1/init", func(r chi.Router) {
		r.Post("/", h.initAccount)
	})

	r.Route("/api/v1/wallet", func(r chi.Router) {
		r.Use(h.validateToken)
		r.Post("/", h.enableWallet)
		r.Get("/", h.getWallet)
		r.Get("/transactions", h.getTransactions)
		r.Post("/deposits", h.deposit)
		r.Post("/withdrawals", h.withdraw)
		r.Patch("/", h.disableWallet)
	})

}
