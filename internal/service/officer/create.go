package officer

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
	"go-sim/internal/pkg/auth"
	"go-sim/pkg/validate"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

// Create implements backend.OfficerServer..
func (s *OfficerService) Create(ctx context.Context, req *v1.CreateReq) (*emptypb.Empty, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, v1.ErrorParamsValidation(err.Error())
	}

	pwd, _ := auth.Encrypt(req.Password)

	data := &biz.Officer{
		Username:  req.Username,
		Password:  pwd,
		Name:      req.Name,
		Mobile:    req.Mobile,
		Status:    1,
		LoginedAt: time.Now(),
		DeletedAt: time.Now(),
	}
	if err := s.uc.Create(ctx, data); err != nil {
		return nil, err
	}
	return nil, nil
}
