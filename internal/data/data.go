package data

import (
	"go-sim/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysql, NewOfficerRepo, NewMenuRepo)

// Data .
type Data struct {
	Module string
	db     *gorm.DB
	log    *log.Helper
}

// NewData .
func NewData(db *gorm.DB, c *conf.Data, logger log.Logger) (*Data, func(), error) {
	module := "go-sim"
	logs := log.NewHelper(log.With(logger, "module ", module))

	data := &Data{
		Module: module,
		db:     db,
	}

	cleanup := func() {
		logs.Info("close go-sim")
	}
	return data, cleanup, nil
}

func NewMysql(conf *conf.Data, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "go-sim"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		logs.Fatalf("mysql connect err : %v", err)
	}
	return db
}
