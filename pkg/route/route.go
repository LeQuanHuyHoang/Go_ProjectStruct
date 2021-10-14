// Định nghĩa các API
package route

import (
	"crawl-data/pkg/handler"
	"crawl-data/pkg/repo"
	srv "crawl-data/pkg/service"

	"github.com/caarlos0/env/v6"
	"gitlab.com/goxp/cloud0/ginext"
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

	db := s.GetDB()

	if s.setting.DbDebugEnable {
		db = db.Debug()
	}

	userRepo := repo.NewRepo(db)

	userService := srv.NewUserService(userRepo)
	authService := srv.NewAuthService(userRepo)
	jwtService := srv.NewJWTService()

	userHandler := handler.NewUserHandler(userService, jwtService)
	authHandler := handler.NewAuthHandler(jwtService, authService)

	//Gom nhom API
	v1Api := s.Router.Group("/api/v1")
	v1Api.POST("/auth/sign-up", ginext.WrapHandler(authHandler.SignUp))
	v1Api.POST("/auth/login", ginext.WrapHandler(authHandler.Login))

	v1Api.GET("/user/info", ginext.WrapHandler(userHandler.Profile))

	migrate := handler.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrate.Migrate)

	return s
}
