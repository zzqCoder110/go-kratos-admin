package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go-sim/internal/biz"
)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m menuRepo) Creates(ctx context.Context, menu *biz.Menu) error {
	return m.data.db.Create(menu).Error
}
