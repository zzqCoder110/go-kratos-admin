package biz

import (
	"context"
	"errors"
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/conf"
	"go-sim/internal/pkg/auth"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.New("用户不存在")
)

// Officer is a Officer model.
type Officer struct {
	Id        int64     `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password,omitempty" gorm:"column:password"`
	Mobile    string    `json:"mobile" gorm:"column:mobile"`
	Email     string    `json:"email" gorm:"column:email"`
	Name      string    `json:"name" gorm:"column:name"`
	Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Status    int64     `json:"status" gorm:"column:status"`
	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:logined_at"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at"`
}

const (
	StatusNormal int = iota + 1
	StatusDelete
)

func (o Officer) TableName() string {
	return "officer"
}

// OfficerRepo is a Greater repo.
type OfficerRepo interface {
	Create(context.Context, *Officer) error
	GetByUsername(ctx context.Context, username string) (*Officer, error)
	Update(context.Context, *Officer) error
	GetByID(context.Context, int64) (*Officer, error)
	List(context.Context, int64, int64, map[string]interface{}) ([]*Officer, int64, error)
}

// OfficerUsecase is a Officer usecase.
type OfficerUsecase struct {
	key  string
	repo OfficerRepo
	log  *log.Helper
}

// NewOfficerUsecase new a Officer usecase.
func NewOfficerUsecase(conf *conf.Jwt, repo OfficerRepo, logger log.Logger) *OfficerUsecase {
	return &OfficerUsecase{key: conf.Key, repo: repo, log: log.NewHelper(logger)}
}

func (uc *OfficerUsecase) Create(ctx context.Context, g *Officer) error {
	return uc.repo.Create(ctx, g)
}

func (uc *OfficerUsecase) Update(ctx context.Context, req *v1.UpdateReq) error {
	officer, err := uc.repo.GetByID(ctx, req.GetId())
	if err != nil {
		return err
	}
	officer.Username = req.Username
	officer.Name = req.Name
	officer.Mobile = req.Mobile
	officer.Status = req.Status
	return uc.repo.Update(ctx, officer)
}

func (uc *OfficerUsecase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRep, error) {
	officer, err := uc.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, v1.ErrorLoginFailed("找不到用户: %s", err.Error())
	}

	if err := auth.Compare(officer.Password, req.Password); err != nil {
		return nil, v1.ErrorLoginFailed("密码不匹配")
	}

	// generate token
	custom := auth.CustomAuth{
		Id:       officer.Id,
		Username: officer.Username,
	}

	token := auth.GenerateJWTToken(custom, []byte(uc.key))
	if err != nil {
		return nil, v1.ErrorLoginFailed("登陆凭证生成失败: %s", err.Error())
	}
	return &v1.LoginRep{
		Token: token,
	}, nil
}

func (uc *OfficerUsecase) List(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*Officer, int64, error) {
	return uc.repo.List(ctx, page, pageSize, params)
}
