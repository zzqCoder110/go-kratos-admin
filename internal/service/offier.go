package service

import (
	"context"

	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
)

// OfficerService is a greeter service.
type OfficerService struct {
	v1.UnimplementedOfficerServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *OfficerService {
	return &OfficerService{uc: uc}
}

// Create implements backend.OfficerServer.
func (s *OfficerService) Create(ctx context.Context, in *v1.CreateReq) (*v1.CreateRep, error) {
	return &v1.CreateRep{}, nil
}
