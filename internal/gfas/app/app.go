package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/phinc27/gfas/config"
	"github.com/phinc27/gfas/internal/gfas"
	"github.com/phinc27/gfas/pkg/es/store"
	"github.com/phinc27/gfas/pkg/eventstroredb"
	"github.com/phinc27/gfas/pkg/logger"
)

type Application struct {
	cfg    *config.Config
	logger logger.Logger

	doneCh chan struct{}
	echo   *echo.Echo

	userAchievementsService *gfas.UserAchievementsService
}

func NewApplication(cfg *config.Config, logger logger.Logger) *Application {
	return &Application{
		cfg:    cfg,
		logger: logger,
		echo:   echo.New(),
		doneCh: make(chan struct{}),
	}
}

func (app *Application) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	db, err := eventstroredb.NewEventStoreDB(app.cfg.EventStoreConfig)
	if err != nil {
		return err
	}
	defer db.Close() // nolint: errcheck

	aggregateStore := store.NewAggregateStore(app.logger, db)

	app.userAchievementsService = gfas.NewUserAchievementsService(app.logger, aggregateStore)

	userAchievementsHandlers := gfas.NewUserAchievementsHandlers(
		app.echo.Group(app.cfg.Http.BasePath).Group(app.cfg.Http.AchievementPath),
		app.logger, app.userAchievementsService,
	)
	userAchievementsHandlers.RegisterRoutes()

	go func() {
		if err := app.runHttpServer(); err != nil {
			app.logger.Errorf("(app.runHttpServer) err: {%v}", err)
			cancel()
		}
	}()

	<-ctx.Done()
	app.logger.Infof("application exited...")
	return nil
}

func (app *Application) runHttpServer() error {
	return app.echo.Start(app.cfg.Http.Addr)
}
