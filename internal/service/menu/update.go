package menu

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
	"go-sim/pkg/validate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MenuService) MenuUpdate(ctx context.Context, req *v1.MenuUpdateReq) (*emptypb.Empty, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, v1.ErrorParamsValidation(err.Error())
	}

	data := &biz.Menu{
		Id:       req.Id,
		Pid:      req.Pid,
		Name:     req.Name,
		Icon:     req.Icon,
		Path:     req.Path,
		Type:     req.Type,
		Sequence: req.Sequence,
	}
	if err := s.uc.Update(ctx, data); err != nil {
		return nil, err
	}
	return nil, nil
}
