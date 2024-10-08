package main

import (
	"context"
	"log"
	"time"

	desc "github.com/igortoigildin/stupefied_bell/pkg/delivery_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address         = "localhost:50051"
	statusDelivered = "delivered"
	statusAccepted  = "accepted"
	orderID         = "1"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewDeliveryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.SetDeliveryStatus(ctx, &desc.SetDeliveryStatusRequest{Status: statusAccepted, OrderId: orderID})
	if err != nil {
		log.Fatalf("failed to update status: %v", err)
	}
}
