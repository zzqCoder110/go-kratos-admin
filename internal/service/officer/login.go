package officer

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/pkg/validate"
)

func (s *OfficerService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRep, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, v1.ErrorParamsValidation(err.Error())
	}

	return s.uc.Login(ctx, req)
}
