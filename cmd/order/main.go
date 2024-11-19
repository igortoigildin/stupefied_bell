package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	config "github.com/igortoigildin/stupefied_bell/config/order"
	"github.com/igortoigildin/stupefied_bell/internal/order/api/grpc/app"
	api "github.com/igortoigildin/stupefied_bell/internal/order/api/rest"
	migrate "github.com/igortoigildin/stupefied_bell/internal/order/migrator"

	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()

	if err := logger.Initialize(cfg.LogLevel); err != nil {
		log.Fatal("error while initializing logger", err)
	}
	logger.Log.Info("starting order app", zap.String("env", cfg.Env))

	// open sql conn and init migration
	db := migrate.New(cfg)

	// grpc
	application := app.New(db, cfg)

	go func() {
		err := application.GRPCServer.MustRun()
		if err != nil {
			logger.Log.Fatal("error while initializing grpc app", zap.Error(err))
		}
	}()

	// кafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers[0],
	})
	if err != nil {
		logger.Log.Info("Kafka failed to initialize", zap.Error(err))
	}
	defer p.Close()

	// rest
	srv := api.New(cfg, db)

	logger.Log.Info("starting http server at:", zap.String("addr:", srv.Addr))

	if err := srv.ListenAndServe(); err != nil {
		logger.Log.Fatal("failed to start server")
	}
}
