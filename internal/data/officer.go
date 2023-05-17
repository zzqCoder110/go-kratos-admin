package data

import (
	"context"

	"go-sim/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type officerRepo struct {
	data *Data
	log  *log.Helper
}

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

func (r *officerRepo) ListByHello(context.Context, string) ([]*biz.Officer, error) {
	return nil, nil
}

func (r *officerRepo) ListAll(context.Context) ([]*biz.Officer, error) {
	return nil, nil
}
