package config

import (
	"github.com/arifmfh/go-mini-wallet/service/delivery/http"
)

func (cfg *Config) initService() error {

	http.Router(cfg.Router)

	return nil
}
