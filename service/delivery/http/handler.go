package http

import (
	"encoding/json"
	"net/http"

	"github.com/arifmfh/go-mini-wallet/models"
	"github.com/go-chi/jwtauth"
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
		Status: "success",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func (h Handler) enableWallet(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	costumerXID := claims["costumer_xid"].(string)

	data, code, err := h.WalletUsecase.EnableWallet(costumerXID)
	if err != nil {
		if code == 400 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.JSONResponse{
				Status: "fail",
				Data: map[string]interface{}{
					"error": err.Error(),
				},
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.JSONResponse{
				Status: "fail",
				Data: map[string]interface{}{
					"error": err.Error(),
				},
			})
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.JSONResponse{
		Status: "success",
		Data: map[string]interface{}{
			"wallet": data,
		},
	})
}

func (h Handler) getWallet(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	costumerXID := claims["costumer_xid"].(string)

	data, code, err := h.WalletUsecase.GetWallet(costumerXID)
	if err != nil {
		if code == 400 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.JSONResponse{
				Status: "fail",
				Data: map[string]interface{}{
					"error": err.Error(),
				},
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.JSONResponse{
				Status: "fail",
				Data: map[string]interface{}{
					"error": err.Error(),
				},
			})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.JSONResponse{
		Status: "success",
		Data: map[string]interface{}{
			"wallet": data,
		},
	})
}