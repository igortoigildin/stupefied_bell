package grpcapp

import (
	"fmt"
	"net"

	delgrpc "github.com/igortoigildin/stupefied_bell/internal/api/delgrpc"
	"github.com/igortoigildin/stupefied_bell/internal/storage/postgres"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpc.Server
	port       int
}

func New(
	port int,
	storage postgres.Repository,
) *App {
	gRPCServer := grpc.NewServer()

	delgrpc.Register(gRPCServer, &storage)

	return &App{
		GRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Log.Info("grpc server is running:", zap.String("addr", l.Addr().String()))

	if err := a.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
