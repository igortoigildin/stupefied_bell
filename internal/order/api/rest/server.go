package rest

import (
	"database/sql"
	"net/http"

	config "github.com/igortoigildin/stupefied_bell/config/order"
	psql "github.com/igortoigildin/stupefied_bell/internal/order/storage/postgres"
)

func New(cfg *config.Config, db *sql.DB) *http.Server {
	storage := psql.NewRepository(db)
	mux := Router(cfg, storage)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      mux,
		ReadTimeout:  cfg.HTTPserver.Timeout,
		WriteTimeout: cfg.HTTPserver.Timeout,
		IdleTimeout:  cfg.HTTPserver.IdleTimout,
	}

	return srv
}
