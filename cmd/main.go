package main

import (
	"one-lab-final/internal/app"
	"one-lab-final/internal/config"
)

// @title           Library API
// @version         0.0.1
// @description     API for Library application

// @contact.name   Kirill Shaforostov
// @contact.email  dragon090986@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.ParseConfig("configs/config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
