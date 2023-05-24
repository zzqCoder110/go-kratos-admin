package officer

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/pkg/validate"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update implements backend.OfficerServer.
func (s *OfficerService) Update(ctx context.Context, req *v1.UpdateReq) (*emptypb.Empty, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, v1.ErrorParamsValidation(err.Error())
	}
	if err := s.uc.Update(ctx, req); err != nil {
		return nil, err
	}
	return nil, nil
}
