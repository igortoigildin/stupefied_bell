package api

import (
	"context"

	"github.com/igortoigildin/stupefied_bell/internal/order/model"
)

//go:generate  mockery --name=OrderRepository
type OrderRepository interface {
	SaveOrder(ctx context.Context, order model.Order) (string, error)
	SelectAllOrders(ctx context.Context) ([]model.Order, error)
	DeleteOrder(ctx context.Context, number string) error
	UpdateOrder(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, orderID string, status string) error
}
