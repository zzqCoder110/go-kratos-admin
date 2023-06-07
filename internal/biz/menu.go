package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Menu is a Menu model.
type Menu struct {
	Id        int64     `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Pid       int64     `json:"pid" gorm:"column:pid"`
	Name      string    `json:"name" gorm:"column:name"`
	Icon      string    `json:"icon" gorm:"column:icon"`
	Path      string    `json:"path" gorm:"column:path"`
	Status    int64     `json:"status" gorm:"column:status"`
	Sequence  int64     `json:"sequence" gorm:"column:sequence"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (m Menu) TableName() string {
	return "menu"
}

type MenuRepo interface {
	Creates(context.Context, *Menu) error
}

type MenuUsecase struct {
	repo MenuRepo
	log  *log.Helper
}

func NewMenuUsecase(repo MenuRepo, logger log.Logger) *MenuUsecase {
	return &MenuUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MenuUsecase) Create(ctx context.Context, m *Menu) error {
	return uc.repo.Creates(ctx, m)
}
