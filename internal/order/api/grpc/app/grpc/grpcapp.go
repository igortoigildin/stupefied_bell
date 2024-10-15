package grpcapp

import (
	"errors"
	"fmt"
	"net"

	delgrpc "github.com/igortoigildin/stupefied_bell/internal/order/api/grpc"
	"github.com/igortoigildin/stupefied_bell/internal/order/storage/postgres"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpc.Server
	port       int
	ip 			net.IP
}

func New(
	port int,
	storage postgres.Repository,
	ip net.IP,
) *App {
	gRPCServer := grpc.NewServer()

	delgrpc.Register(gRPCServer, &storage)

	return &App{
		GRPCServer: gRPCServer,
		port:       port,
		ip:			ip,
	}
}

func (a *App) MustRun() error {
	if err := a.Run(); err != nil {
		logger.Log.Error("failed to run grpc app", zap.Error(err))
		return errors.New("failed to start grpc app")
	}
	return nil
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
