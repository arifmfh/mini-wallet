package config

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (c *Config) initChi() error {
	c.Router = chi.NewRouter()
	c.Router.Use(middleware.RequestID)
	c.Router.Use(middleware.RealIP)
	c.Router.Use(middleware.Logger)
	c.Router.Use(middleware.Recoverer)
	c.Router.Use(middleware.Timeout(60 * time.Second))
	c.Router.Use(middleware.SetHeader("Content-Type", "application/json"))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Encoding", "Cookie", "Origin", "X-Api-Key"},
	})
	c.Router.Use(cors.Handler)

	return nil
}
