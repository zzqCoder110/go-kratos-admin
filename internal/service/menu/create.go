package menu

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
	"go-sim/pkg/validate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MenuService) MenuCreate(ctx context.Context, req *v1.MenuCreateReq) (*emptypb.Empty, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, v1.ErrorParamsValidation(err.Error())
	}

	data := &biz.Menu{
		Pid:      req.Pid,
		Name:     req.Name,
		Icon:     req.Icon,
		Path:     req.Path,
		Sequence: req.Sequence,
		Status:   1,
	}
	if err := s.uc.Create(ctx, data); err != nil {
		return nil, err
	}
	return nil, nil
}
