package gfas

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/hiendaovinh/toolkit/pkg/errorx"
	"github.com/hiendaovinh/toolkit/pkg/httpx-echo"
	"github.com/labstack/echo/v4"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/pkg/errors"
)

type UserAchievementsHandlers interface {
	ClaimUserAchievement() echo.HandlerFunc

	GetUserAchievementsByID() echo.HandlerFunc

	RegisterRoutes()
}

type userAchievementsHandlers struct {
	group   *echo.Group
	logger  logger.Logger
	service *UserAchievementsService
}

func NewUserAchievementsHandlers(group *echo.Group, logger logger.Logger, service *UserAchievementsService) UserAchievementsHandlers {
	return &userAchievementsHandlers{
		group:   group,
		logger:  logger,
		service: service,
	}
}

func (handlers *userAchievementsHandlers) ClaimUserAchievement() echo.HandlerFunc {
	return func(c echo.Context) error {
		// FIXME: consider get userID from context when authentication is implemented
		var request struct {
			UserID        string `json:"user_id"`
			AchievementID string `json:"achievement_id"`
			Tier          string `json:"tier"`
		}

		err := c.Bind(&request)
		if err != nil {
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Validation))
		}

		command := NewClaimUserAchievementCommand(request.UserID, AchievementID(request.AchievementID), AchievementTier(request.Tier))
		err = handlers.service.commands.claimUserAchievement.Handle(c.Request().Context(), command)
		if err != nil {
			if errors.Is(err, esdb.ErrStreamNotFound) {
				return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Validation))
			}
			if errors.Is(err, ErrAchievementNotFound) || errors.Is(err, ErrAchievementNotAchievedYet) {
				return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Validation))
			}
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Service))
		}

		return httpx.RestAbort(c, nil, nil)
	}
}

func (handlers *userAchievementsHandlers) GetUserAchievementsByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		query := NewGetUserAchievementsByIDQuery(id)
		projection, err := handlers.service.queries.UserAchievementsByID.Handle(c.Request().Context(), query)
		if err != nil {
			return httpx.RestAbort(c, nil, errorx.Wrap(err, errorx.Service))
		}

		return httpx.RestAbort(c, projection, nil)
	}
}

func (handlers *userAchievementsHandlers) RegisterRoutes() {
	handlers.group.GET("/by/user-id/:id", handlers.GetUserAchievementsByID())
	handlers.group.POST("/claim", handlers.ClaimUserAchievement())
}
