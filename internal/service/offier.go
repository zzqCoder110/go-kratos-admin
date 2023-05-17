package service

import (
	"context"
	"go-sim/internal/pkg/auth"

	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
)

// OfficerService is a greeter service.
type OfficerService struct {
	v1.UnimplementedOfficerServer

	uc *biz.OfficerUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.OfficerUsecase) *OfficerService {
	return &OfficerService{uc: uc}
}

// Create implements backend.OfficerServer.
func (s *OfficerService) Create(ctx context.Context, in *v1.CreateReq) (*v1.CreateRep, error) {
	pwd, _ := auth.Encrypt(in.Password)
	data := &biz.Officer{
		Username: in.Username,
		Password: pwd,
		Name:     in.Name,
		Mobile:   in.Mobile,
		Status:   1,
	}
	if err := s.uc.Create(ctx, data); err != nil {
		return nil, err
	}
	return &v1.CreateRep{}, nil
}
