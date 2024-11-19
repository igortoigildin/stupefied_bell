package delgrpc

import (
	"context"
	"errors"

	model "github.com/igortoigildin/stupefied_bell/internal/order/model"
	delivery "github.com/igortoigildin/stupefied_bell/pkg/delivery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate  mockery --name=OrderRepository
type OrderRepository interface {
	SaveOrder(ctx context.Context, order model.Order) (string, error)
	SelectAllOrders(ctx context.Context) ([]model.Order, error)
	DeleteOrder(ctx context.Context, number string) error
	UpdateOrder(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

type ServerAPI struct {
	delivery.UnimplementedDeliveryServiceServer
	OrderRepository OrderRepository
}

func Register(gRPC *grpc.Server, repo OrderRepository) {
	delivery.RegisterDeliveryServiceServer(gRPC, &ServerAPI{OrderRepository: repo})
}

func (s *ServerAPI) SetDeliveryStatus(ctx context.Context, req *delivery.SetStatusRequest) (*emptypb.Empty, error) {
	err := s.OrderRepository.UpdateStatus(ctx, req.GetOrderId(), req.GetStatus().Enum().String())
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return nil, status.Error(codes.InvalidArgument, "invalid order_id")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return nil, nil
}
