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

func (r *officerRepo) Update(ctx context.Context, officer *biz.Officer) error {
	return r.data.db.Save(officer).Error
}

func (r *officerRepo) GetByID(ctx context.Context, id int64) (*biz.Officer, error) {
	data := &biz.Officer{}
	err := r.data.db.First(data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz.ErrUserNotFound
		}
		return nil, err
	}
	return data, nil
}

func (r *officerRepo) GetByUsername(ctx context.Context, username string) (*biz.Officer, error) {
	var data = &biz.Officer{}
	err := r.data.db.Where("name = ? and status = ?", username, biz.StatusNormal).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz.ErrUserNotFound
		}
		return nil, err
	}

	return data, nil
}

func (r *officerRepo) GetById(ctx context.Context, id int64) (*biz.Officer, error) {
	var data = &biz.Officer{}
	err := r.data.db.First(data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz.ErrUserNotFound
		}
		return nil, err
	}

	return data, nil
}

func (r *officerRepo) List(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*biz.Officer, int64, error) {
	query := r.data.db
	if params["username"].(string) != "" {
		query = query.Where("username = ?", params["username"])
	}
	if params["name"].(string) != "" {
		query = query.Where("name LIKE ?", "%"+params["name"].(string)+"%")
	}
	if params["mobile"].(string) != "" {
		query = query.Where("mobile = ?", params["mobile"])
	}
	if params["status"].(int64) != 0 {
		query = query.Where("status = ?", params["status"])
	}

	var data []*biz.Officer
	count := int64(0)
	query.Find(&data).Count(&count)
	return data, count, nil
}
