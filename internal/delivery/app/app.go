package app

import (
	"log"
	"net/http"

	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	api "github.com/igortoigildin/stupefied_bell/internal/delivery/api/rest"
	"github.com/igortoigildin/stupefied_bell/pkg/kafka"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type (
	orderContext struct {
		echo.Context
		status string
		id     int
	}
)

func Run(cfg *config.Config) {

	// Kafka consumer
	cons, err := kafka.NewConsumer(cfg)
	if err != nil {
		logger.Log.Fatal("failed to initialize consumer: %v", zap.Error(err))
	}
	defer cons.Close()
	cons.SubscribeTopics([]string{cfg.Kafka.Topic}, nil)

	go kafka.RunReading(cons)

	// HTTP Server
	e := echo.New()
	api.NewRouter(e, cfg)

	if err := e.Start(cfg.HTTPserver.Address); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
