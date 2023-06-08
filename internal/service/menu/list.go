package menu

import (
	"context"
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MenuService) MenuList(ctx context.Context, req *emptypb.Empty) (*v1.MenuGetListRep, error) {
	menus, err := s.uc.GetList(ctx)
	if err != nil {
		return nil, err
	}

	list := convertToMenuGetListRep(menus)

	return &v1.MenuGetListRep{
		List: list,
	}, nil
}

func convertToMenuGetListRep(menus []*biz.Menu) []*v1.MenuInfo {
	var v1Menus []*v1.MenuInfo

	for _, menu := range menus {
		v1Menu := &v1.MenuInfo{
			Id:       menu.Id,
			Pid:      menu.Pid,
			Name:     menu.Name,
			Icon:     menu.Icon,
			Path:     menu.Path,
			Type:     menu.Type,
			Sequence: menu.Sequence,
		}

		if len(menu.Children) > 0 {
			v1Menu.Children = convertToMenuGetListRep(menu.Children)
		}

		v1Menus = append(v1Menus, v1Menu)
	}

	return v1Menus
}
