package app

import (
	"database/sql"

	config "github.com/igortoigildin/stupefied_bell/config/order"
	grpcapp "github.com/igortoigildin/stupefied_bell/internal/order/api/grpc/app/grpc"
	"github.com/igortoigildin/stupefied_bell/internal/order/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	db *sql.DB,
	config *config.Config,
) *App {
	storage := postgres.NewRepository(db)

	grpcApp := grpcapp.New(config.Port, *storage, config.Ip)

	return &App{
		GRPCServer: grpcApp,
	}
}
