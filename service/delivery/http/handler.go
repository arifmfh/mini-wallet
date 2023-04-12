package http

import (
	"encoding/json"
	"net/http"

	"github.com/arifmfh/go-mini-wallet/models"
)

func (h Handler) initAccount(w http.ResponseWriter, r *http.Request) {
	costumerXID := r.FormValue("customer_xid")
	if costumerXID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.JSONResponse{
			Status: "fail",
			Data: map[string]interface{}{
				"error": map[string]interface{}{
					"customer_xid": []string{"Missing data for required field."},
				},
			},
		})
		return
	}

	token, err := h.claimToken(costumerXID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.JSONResponse{
			Status: "fail",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	err = h.WalletUsecase.Register(costumerXID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.JSONResponse{
			Status: "fail",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.JSONResponse{
		Status: "fail",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func (h Handler) enableWallet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("ok"))
}
