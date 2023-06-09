package menu

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/pkg/validate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MenuService) MenuSort(ctx context.Context, req *v1.MenuSortReq) (*emptypb.Empty, error) {
	if err := validate.ValidateStructCN(req); err != nil {
		return nil, v1.ErrorParamsValidation(err.Error())
	}

	if err := s.uc.SortNodes(ctx, req); err != nil {
		return nil, err
	}
	return nil, nil
}
