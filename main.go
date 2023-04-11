package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arifmfh/go-mini-wallet/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	cfg := config.Init()

	fmt.Println(`Server is running at ` + appHost + appPort)

	err = http.ListenAndServe(appPort, cfg.Router)
	if err != nil {
		log.Fatal(err)
	}
}
