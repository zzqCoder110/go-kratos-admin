package service

import (
	"context"
	"errors"
	"go-sim/internal/pkg/auth"
	"go-sim/pkg/validate"

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
func (s *OfficerService) Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRep, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, errors.New(err.Error())
	}
	return &v1.CreateRep{}, nil
	pwd, _ := auth.Encrypt(req.Password)
	data := &biz.Officer{
		Username: req.Username,
		Password: pwd,
		Name:     req.Name,
		Mobile:   req.Mobile,
		Status:   1,
	}
	if err := s.uc.Create(ctx, data); err != nil {
		return nil, err
	}
	return &v1.CreateRep{}, nil
}
