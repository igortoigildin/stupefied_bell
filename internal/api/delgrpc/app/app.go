package app

import (
	"database/sql"

	grpcapp "github.com/igortoigildin/stupefied_bell/internal/api/delgrpc/app/grpc"
	"github.com/igortoigildin/stupefied_bell/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	db *sql.DB,
	grpcPort int,
) *App {
	storage := postgres.NewRepository(db)

	grpcApp := grpcapp.New(grpcPort, *storage)

	return &App{
		GRPCServer: grpcApp,
	}
}
