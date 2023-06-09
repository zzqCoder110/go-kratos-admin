package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "go-sim/api/backend/v1"
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

	m.data.db.Order("sequence asc").Find(&data)

	return data, nil
}

func (m *menuRepo) SortNodes(ctx context.Context, req *v1.MenuSortReq) error {
	return m.data.db.Transaction(func(tx *gorm.DB) error {
		return appendSortMenuNode(tx, req.List, 0)
	})
}

func appendSortMenuNode(tx *gorm.DB, nodes []*v1.MenuSortReqChildren, parentId int64) error {
	for index, node := range nodes {
		if node.Children != nil {
			if err := appendSortMenuNode(tx, node.Children, node.Id); err != nil {
				return err
			}
		}
		if err := tx.Model(&biz.Menu{}).Where("id = ?", node.GetId()).Updates(map[string]interface{}{
			"pid":      parentId,
			"sequence": index + 1,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
