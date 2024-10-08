package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/igortoigildin/stupefied_bell/internal/config"
	"github.com/igortoigildin/stupefied_bell/internal/model"
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
	// gRPC
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

	// Kafka consumer
	cfg := config.MustLoad()

	cons, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers[0],
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("failed to initialize consumer: %v", err)
	}
	defer cons.Close()

	cons.SubscribeTopics([]string{cfg.Kafka.Topic}, nil)

	for {
		msg, err := cons.ReadMessage(-1)
		if err == nil {
			var order model.Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			fmt.Printf("Received Order: %+v\n", order)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
