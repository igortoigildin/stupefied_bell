package main

import (
	"log"

	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	"github.com/igortoigildin/stupefied_bell/internal/delivery/app"
	logger "github.com/igortoigildin/stupefied_bell/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	if err := logger.Initialize(cfg.LogLevel); err != nil {
		log.Fatal("error while initializing logger", err)
	}

	// Delivery app
	app.Run(cfg)
}
