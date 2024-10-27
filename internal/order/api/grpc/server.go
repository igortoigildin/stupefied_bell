package delgrpc

import (
	"context"
	"errors"

	api "github.com/igortoigildin/stupefied_bell/internal/order/api"
	storage "github.com/igortoigildin/stupefied_bell/internal/order/storage"
	delivery "github.com/igortoigildin/stupefied_bell/pkg/delivery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	delivery.UnimplementedDeliveryServiceServer
	OrderRepository api.OrderRepository
}

func Register(gRPC *grpc.Server, repo api.OrderRepository) {
	delivery.RegisterDeliveryServiceServer(gRPC, &ServerAPI{OrderRepository: repo})
}

func (s *ServerAPI) SetDeliveryStatus(ctx context.Context, req *delivery.SetStatusRequest) (*emptypb.Empty, error) {
	err := s.OrderRepository.UpdateStatus(ctx, req.GetOrderId(), req.GetStatus().Enum().String())
	if err != nil {
		if errors.Is(err, storage.ErrOrderNotFound) {
			return nil, status.Error(codes.InvalidArgument, "invalid order_id")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return nil, nil
}
