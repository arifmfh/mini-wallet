package models

import "github.com/dgrijalva/jwt-go"

type (
	JSONResponse struct {
		Status interface{} `json:"status"`
		Data   interface{} `json:"data"`
	}

	ClaimJWT struct {
		CostumerXID string `json:"costumer_xid"`
		StdJWT      jwt.StandardClaims
	}
)
