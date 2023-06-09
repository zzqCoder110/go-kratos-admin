package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "go-sim/api/backend/v1"
	"time"
)

var (
	ErrNotFound = errors.New("未找到内容")
)

// Menu is a Menu model.
type Menu struct {
	Id        int64     `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Pid       int64     `json:"pid" gorm:"column:pid"`
	Name      string    `json:"name" gorm:"column:name"`
	Icon      string    `json:"icon" gorm:"column:icon"`
	Path      string    `json:"path" gorm:"column:path"`
	Type      int64     `json:"type" gorm:"column:type"`
	Sequence  int64     `json:"sequence" gorm:"column:sequence"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	Children  []*Menu   `gorm:"-"`
}

const (
	TypeDir int = iota + 1
	TypePage
	TypeButton
)

func (m Menu) TableName() string {
	return "menu"
}

type MenuRepo interface {
	Create(context.Context, *Menu) error
	Update(context.Context, *Menu) error
	Delete(context.Context, int64) error
	GetById(context.Context, int64) (*Menu, error)
	GetList(context.Context, map[string]interface{}) ([]*Menu, error)
	SortNodes(context.Context, *v1.MenuSortReq) error
}

type MenuUsecase struct {
	repo MenuRepo
	log  *log.Helper
}

func NewMenuUsecase(repo MenuRepo, logger log.Logger) *MenuUsecase {
	return &MenuUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MenuUsecase) Create(ctx context.Context, m *Menu) error {
	return uc.repo.Create(ctx, m)
}

func (uc *MenuUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *MenuUsecase) Update(ctx context.Context, m *Menu) error {
	menu, err := uc.repo.GetById(ctx, m.Id)
	if err != nil {
		return err
	}
	m.CreatedAt = menu.CreatedAt
	return uc.repo.Update(ctx, m)
}

func (uc *MenuUsecase) GetListTree(ctx context.Context) ([]*Menu, error) {
	menus, err := uc.repo.GetList(ctx, nil)
	if err != nil {
		return nil, err
	}
	list := toTree(menus, 0)
	return list, nil
}

func (uc *MenuUsecase) SortNodes(ctx context.Context, req *v1.MenuSortReq) error {
	return uc.repo.SortNodes(ctx, req)
}

func toTree(menus []*Menu, parentId int64) []*Menu {
	var subMenus []*Menu

	for _, menu := range menus {
		if menu.Pid == parentId {
			children := toTree(menus, menu.Id)
			menu.Children = children
			subMenus = append(subMenus, menu)
		}
	}
	return subMenus
}
