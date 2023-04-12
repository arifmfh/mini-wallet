package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/arifmfh/go-mini-wallet/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

type ClaimJWT struct {
	CostumerXID string `json:"costumer_xid"`
	jwt.StandardClaims
}

func (h Handler) claimToken(costumerXID string) (string, error) {
	claims := ClaimJWT{
		costumerXID,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("APP_ENV") + os.Getenv("JWT_KEY") + os.Getenv("APP_NAME")))
}

func (h Handler) validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.JSONResponse{
				Status: "fail",
				Data: map[string]string{
					"error": fmt.Sprintf("Bearer Token - %v", err),
				},
			})
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
