package gfas

import (
	"github.com/phinc275/gfas/pkg/es"
	"github.com/phinc275/gfas/pkg/logger"
)

type UserAchievementsService struct {
	commands *UserAchievementsCommands
	queries  *UserAchievementsQueries
}

func NewUserAchievementsService(
	logger logger.Logger,
	es es.AggregateStore,
) *UserAchievementsService {
	applyExternalHandler := NewApplyExternalEventCommandHandler(logger, es)
	claimUserAchievementHandler := NewClaimUserAchievementCommandHandler(logger, es)

	commands := NewUserAchievementsCommands(
		applyExternalHandler,
		claimUserAchievementHandler,
	)

	userAchievementsByIDHandler := NewUserAchievementsByIDHandler(logger, es)
	userPublicAchievementsByIDHandler := NewUserPublicAchievementsByIDHandler(logger, es)
	queries := NewUserAchievementsQueries(userAchievementsByIDHandler, userPublicAchievementsByIDHandler)

	return &UserAchievementsService{
		commands: commands,
		queries:  queries,
	}
}
