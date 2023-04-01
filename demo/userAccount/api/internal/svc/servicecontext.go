package svc

import (
	"demo/userAccount/api/internal/config"
	"demo/userAccount/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel models.UserTableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: models.NewUserTableModel(sqlx.NewMysql(c.DataSource)),
	}
}
