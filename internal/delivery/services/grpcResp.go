package services

import (
	"context"

	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	delivery "github.com/igortoigildin/stupefied_bell/pkg/delivery_v1"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	statusDelivered = "delivered"
	statusAccepted  = "accepted"
)

func SendGRPCResponse(cfg *config.Config, status string, id string) error {
	conn, err := grpc.Dial(cfg.GRPCServer.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Error("failed to connect to server:", zap.Error(err))
		return err
	}
	defer conn.Close()

	c := delivery.NewDeliveryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GRPCServer.Timeout)
	defer cancel()
	
	if status == statusDelivered {
		_, err = c.SetDeliveryStatus(ctx, &delivery.SetDeliveryStatusRequest{Status: 1, OrderId: id}) 
		if err != nil {
			logger.Log.Error("failed to update status:", zap.Error(err))
			return err
		}
	} else if status == statusAccepted {
		_, err = c.SetDeliveryStatus(ctx, &delivery.SetDeliveryStatusRequest{Status: 2, OrderId: id})
		if err != nil {
			logger.Log.Error("failed to update status:", zap.Error(err))
			return err
		}
	}

	return nil
}
