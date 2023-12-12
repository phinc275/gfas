package gfas

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/hiendaovinh/toolkit/pkg/auth"
	"github.com/hiendaovinh/toolkit/pkg/errorx"
	"github.com/hiendaovinh/toolkit/pkg/httpx-echo"
	"github.com/labstack/echo/v4"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/pkg/errors"
)

type UserAchievementsHandlers interface {
	ClaimUserAchievement() echo.HandlerFunc
	GetUserAchievements() echo.HandlerFunc
	RegisterRoutes()
}

type userAchievementsHandlers struct {
	group           *echo.Group
	authMiddlewares []echo.MiddlewareFunc
	logger          logger.Logger
	service         *UserAchievementsService
}

func NewUserAchievementsHandlers(group *echo.Group, authMiddlewares []echo.MiddlewareFunc, logger logger.Logger, service *UserAchievementsService) UserAchievementsHandlers {
	return &userAchievementsHandlers{
		group:           group,
		authMiddlewares: authMiddlewares,
		logger:          logger,
		service:         service,
	}
}

func (handlers *userAchievementsHandlers) ClaimUserAchievement() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := auth.ResolveValidSubject(ctx)
		if err != nil {
			return httpx.RestAbort(c, nil, err)
		}

		var request struct {
			AchievementID string `json:"achievement_id"`
			Tier          string `json:"tier"`
		}

		err = c.Bind(&request)
		if err != nil {
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Validation))
		}

		command := NewClaimUserAchievementCommand(id, AchievementID(request.AchievementID), AchievementTier(request.Tier))
		err = handlers.service.commands.claimUserAchievement.Handle(c.Request().Context(), command)
		if err != nil {
			if errors.Is(err, esdb.ErrStreamNotFound) {
				return httpx.RestAbort(c, nil, errorx.Wrap(ErrInvalidUserAchievement, errorx.Validation))
			}
			if errors.Is(err, ErrAchievementNotFound) || errors.Is(err, ErrAchievementNotAchievedYet) || errors.Is(err, ErrAchievementAlreadyClaimed) {
				return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Validation))
			}
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Service))
		}

		return httpx.RestAbort(c, nil, nil)
	}
}

func (handlers *userAchievementsHandlers) GetUserAchievements() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id, err := auth.ResolveValidSubject(ctx)
		if err != nil {
			return httpx.RestAbort(c, nil, err)
		}

		query := NewGetUserAchievementsByIDQuery(id)
		projection, err := handlers.service.queries.UserAchievementsByID.Handle(c.Request().Context(), query)
		if err != nil {
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Service))
		}

		return httpx.RestAbort(c, projection, nil)
	}
}

func (handlers *userAchievementsHandlers) GetUserPublicAchievements() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("userID")
		query := NewGetUserPublicAchievementsByIDQuery(id)
		projection, err := handlers.service.queries.UserPublicAchievementsByID.Handle(c.Request().Context(), query)
		if err != nil {
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Service))
		}

		return httpx.RestAbort(c, projection, nil)
	}
}

func (handlers *userAchievementsHandlers) RegisterRoutes() {
	handlers.group.GET("/", handlers.GetUserAchievements(), handlers.authMiddlewares...)
	handlers.group.POST("/claim", handlers.ClaimUserAchievement(), handlers.authMiddlewares...)
	handlers.group.GET("/public/:userID", handlers.GetUserPublicAchievements())
}
