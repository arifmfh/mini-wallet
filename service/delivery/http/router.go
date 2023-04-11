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

	r.Get("/", h.welcome)
}
