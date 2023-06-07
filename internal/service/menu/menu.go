package menu

import (
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
)

type MenuService struct {
	v1.UnimplementedMenuServer

	uc *biz.MenuUsecase
}

// NewMenuService new a greeter service.
func NewMenuService(uc *biz.MenuUsecase) *MenuService {
	return &MenuService{uc: uc}
}
