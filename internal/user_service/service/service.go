package service

import (
	"context"
	"github.com/EridanSilver/clay-test/pkg/pb/user"
	"github.com/utrack/clay/v2/transport"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetById(ctx context.Context, req *user.GetByIdRequest) (*user.GetByIdResponse, error) {
	return &user.GetByIdResponse{
		Data: "SomeData",
	}, nil
}

func (s *Service) GetDescription() transport.ServiceDesc {
	return user.NewUserServiceServiceDesc(s)
}
