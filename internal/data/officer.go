package data

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"go-sim/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type officerRepo struct {
	data *Data
	log  *log.Helper
}

const (
	StatusNormal int = iota + 1
	StatusDelete
)

// NewOfficerRepo .
func NewOfficerRepo(data *Data, logger log.Logger) biz.OfficerRepo {
	return &officerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *officerRepo) Create(ctx context.Context, g *biz.Officer) error {
	return r.data.db.Create(g).Error
}

func (r *officerRepo) Update(ctx context.Context, g *biz.Officer) (*biz.Officer, error) {
	return g, nil
}

func (r *officerRepo) FindByID(context.Context, int64) (*biz.Officer, error) {
	return nil, nil
}

func (r *officerRepo) GetByUsername(ctx context.Context, username string) (*biz.Officer, error) {
	var data = &biz.Officer{}
	err := r.data.db.Where("name = ? and status = ?", username, StatusNormal).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz.ErrUserNotFound
		}

		return nil, err
	}

	return data, nil
}

func (r *officerRepo) ListByHello(context.Context, string) ([]*biz.Officer, error) {
	return nil, nil
}

func (r *officerRepo) ListAll(context.Context) ([]*biz.Officer, error) {
	return nil, nil
}
