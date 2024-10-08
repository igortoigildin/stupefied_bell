package delgrpc

import (
	"context"
	"errors"
	"fmt"

	api "github.com/igortoigildin/stupefied_bell/internal/api/http"
	"github.com/igortoigildin/stupefied_bell/internal/storage"
	deliveryV1 "github.com/igortoigildin/stupefied_bell/pkg/delivery_v1"
	desc "github.com/igortoigildin/stupefied_bell/pkg/delivery_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	desc.UnimplementedDeliveryServer
	OrderRepository api.OrderRepository
}

func Register(gRPC *grpc.Server, repo api.OrderRepository) {
	deliveryV1.RegisterDeliveryServer(gRPC, &ServerAPI{OrderRepository: repo})
}

func (s *ServerAPI) SetDeliveryStatus(ctx context.Context, req *desc.SetDeliveryStatusRequest) (*emptypb.Empty, error) {
	if req == nil {
		fmt.Println(req)
		return nil, nil
	}

	err := s.OrderRepository.UpdateStatus(ctx, req.GetOrderId(), req.GetStatus())
	if err != nil {
		if errors.Is(err, storage.ErrOrderNotFound) {
			return nil, status.Error(codes.InvalidArgument, "invalid order_id")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return nil, nil
}
