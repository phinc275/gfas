package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/phinc275/gfas/config"
	"github.com/phinc275/gfas/internal/gfas/app"
	"github.com/phinc275/gfas/pkg/logger"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName(cfg.AppName)
	appLogger.Fatal(app.NewApplication(cfg, appLogger).Run())
}
