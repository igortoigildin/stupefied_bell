package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	api "github.com/igortoigildin/stupefied_bell/internal/api"
	"github.com/igortoigildin/stupefied_bell/internal/config"
	"github.com/igortoigildin/stupefied_bell/internal/logger"
	psql "github.com/igortoigildin/stupefied_bell/internal/storage/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()

	if err := logger.Initialize(cfg.LogLevel); err != nil {
		logger.Log.Info("error while initializing logger", zap.Error(err))
	}
	logger.Log.Info("starting ecommerce app", zap.String("env", cfg.Env))

	db, err := sql.Open("pgx", cfg.DBURI)
	if err != nil {
		logger.Log.Fatal("error while connectiong to DB", zap.Error(err))
	}
	logger.Log.Info("database connection pool established")

	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Log.Fatal("migration error", zap.Error(err))
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", instance)
	if err != nil {
		logger.Log.Fatal("migration error", zap.Error(err))
	}
	err = migrator.Up()
	if err != nil || errors.Is(err, migrate.ErrNoChange) {
		logger.Log.Fatal("migration error", zap.Error(err))
	}
	logger.Log.Info("database connection established")

	storage := psql.NewRepository(db)
	if err != nil {
		logger.Log.Fatal("failed to init storage", zap.Error(err))
	}

	mux := api.Router(cfg, storage)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      mux,
		ReadTimeout:  cfg.HTTPserver.Timeout,
		WriteTimeout: cfg.HTTPserver.Timeout,
		IdleTimeout:  cfg.HTTPserver.IdleTimout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Log.Fatal("failed to start server")
	}
}
