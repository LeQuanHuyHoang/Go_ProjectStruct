// Định nghĩa các API
package route

import (
	"github.com/caarlos0/env/v6"
	"gitlab.com/goxp/cloud0/service"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting *extraSetting
}

func NewService() *Service {
	s := &Service{
		service.NewApp("Crawl data Service", "v1.0"),
		&extraSetting{},
	}

	_ = env.Parse(s.setting)

	/* 	db := s.GetDB()

	   	if s.setting.DbDebugEnable {
	   		db = db.Debug()
	   	} */

	//Gom nhom API
	v1Api := s.Router.Group("/api/v1")
	v1Api.POST("/user/sign-up", nil)
	v1Api.POST("/auth/login", nil)

	return s
}
