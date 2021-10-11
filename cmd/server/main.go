package main

import (
	"context"
	"crawl-data/conf"
	"crawl-data/pkg/route"
	"os"

	"gitlab.com/goxp/cloud0/logger"
)

const (
	APPNAME = "Crawl"
)

func main() {
	conf.SetEnv()

	_ = os.Setenv("PORT", conf.LoadEnv().Port)

	_ = os.Setenv("DB_HOST", conf.LoadEnv().DBHost)
	_ = os.Setenv("DB_PORT", conf.LoadEnv().DBPort)
	_ = os.Setenv("DB_USER", conf.LoadEnv().DBUser)
	_ = os.Setenv("DB_PASS", conf.LoadEnv().DBPass)
	_ = os.Setenv("DB_NAME", conf.LoadEnv().DBName)
	_ = os.Setenv("ENABLE_DB", conf.LoadEnv().EnableDB)

	logger.Init(APPNAME)

	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}

	os.Clearenv()
}
