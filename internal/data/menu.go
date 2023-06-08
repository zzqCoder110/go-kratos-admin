package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"go-sim/internal/biz"
	"gorm.io/gorm"
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

func (m *menuRepo) Create(ctx context.Context, menu *biz.Menu) error {
	return m.data.db.Create(menu).Error
}

func (m *menuRepo) Delete(ctx context.Context, id int64) error {
	return m.data.db.Delete(&biz.Menu{}, id).Error
}

func (m *menuRepo) Update(ctx context.Context, menu *biz.Menu) error {
	return m.data.db.Save(menu).Error
}

func (m *menuRepo) GetById(ctx context.Context, id int64) (*biz.Menu, error) {
	data := &biz.Menu{}
	err := m.data.db.First(data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz.ErrNotFound
		}
		return nil, err
	}
	return data, nil
}

func (m *menuRepo) GetList(ctx context.Context, params map[string]interface{}) ([]*biz.Menu, error) {
	var data []*biz.Menu

	m.data.db.Find(&data)

	return data, nil
}
