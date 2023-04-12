package models

type (
	Wallet struct {
		ID        string  `json:"id"`
		OwnedBy   string  `json:"owned_by"`
		Status    string  `json:"status"`
		EnabledAt string  `json:"enabled_at"`
		Balance   float64 `json:"balance"`
	}

	Transaction struct {
		ID           string  `json:"id"`
		Status       string  `json:"status"`
		TransactedAt string  `json:"transacted_at"`
		Type         string  `json:"type"`
		Amount       float64 `json:"amount"`
		ReferenceID  string  `json:"reference_id"`
	}

	Deposit struct {
		ID          string  `json:"id"`
		DepositedBy string  `json:"deposited_by"`
		Status      string  `json:"status"`
		DepositedAt string  `json:"deposited_at"`
		Amount      float64 `json:"amount"`
		ReferenceID string  `json:"reference_id"`
	}

	Withdraw struct {
		ID          string  `json:"id"`
		WithdrawnBy string  `json:"withdrawn_by"`
		Status      string  `json:"status"`
		WithdrawnAt string  `json:"withdrawn_at"`
		Amount      float64 `json:"amount"`
		ReferenceID string  `json:"reference_id"`
	}
)
