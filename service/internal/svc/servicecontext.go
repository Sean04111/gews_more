package svc

import (
	"gews_more/service/internal/config"
	"gews_more/service/model"
	"gews_more/service/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
	SnapModel model.SnapsModel
	Logincheck rest.Middleware
	Getcode rest.Middleware
	Registercheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		SnapModel: model.NewSnapsModel(sqlx.NewMysql(c.DB.DataSource)),
		Logincheck: middleware.NewLogincheckMiddleware().Handle,
		Getcode:middleware.NewGetcodeMiddleware().Handle,
		Registercheck: middleware.NewRegistercheckMiddleware().Handle,
	}
}
