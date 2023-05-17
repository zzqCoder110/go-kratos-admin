package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
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
	LoginedAt time.Time `json:"loginedAt" gorm:"column:logined_at"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at"`
}

func (o Officer) TableName() string {
	return "officer"
}

// OfficerRepo is a Greater repo.
type OfficerRepo interface {
	Create(context.Context, *Officer) error
	Update(context.Context, *Officer) (*Officer, error)
	FindByID(context.Context, int64) (*Officer, error)
	ListAll(context.Context) ([]*Officer, error)
}

// OfficerUsecase is a Officer usecase.
type OfficerUsecase struct {
	repo OfficerRepo
	log  *log.Helper
}

// NewOfficerUsecase new a Officer usecase.
func NewOfficerUsecase(repo OfficerRepo, logger log.Logger) *OfficerUsecase {
	return &OfficerUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OfficerUsecase) Create(ctx context.Context, g *Officer) error {
	//uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Create(ctx, g)
}
