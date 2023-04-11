package config

import (
	"github.com/go-chi/chi"
)

type Config struct {
	Router *chi.Mux
}

func Init() *Config {
	var cfg Config

	err := cfg.initChi()
	if err != nil {
		panic(err.Error())
	}

	err = cfg.initService()
	if err != nil {
		panic(err.Error())
	}

	return &cfg
}
