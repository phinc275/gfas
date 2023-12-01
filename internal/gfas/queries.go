package gfas

import (
	"context"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/phinc275/gfas/pkg/es"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/pkg/errors"
)

type UserAchievementsQueries struct {
	UserAchievementsByID       UserAchievementsByIDHandler
	UserPublicAchievementsByID UserPublicAchievementsByIDHandler
}

func NewUserAchievementsQueries(
	userAchievementsByID UserAchievementsByIDHandler,
	userPublicAchievementsByID UserPublicAchievementsByIDHandler,
) *UserAchievementsQueries {
	return &UserAchievementsQueries{UserAchievementsByID: userAchievementsByID, UserPublicAchievementsByID: userPublicAchievementsByID}
}

type GetUserAchievementsByIDQuery struct {
	ID string
}

func NewGetUserAchievementsByIDQuery(id string) *GetUserAchievementsByIDQuery {
	return &GetUserAchievementsByIDQuery{ID: id}
}

type UserAchievementsByIDHandler interface {
	Handle(ctx context.Context, query *GetUserAchievementsByIDQuery) ([]AchievementsProjection, error)
}

func NewUserAchievementsByIDHandler(logger logger.Logger, es es.AggregateStore) UserAchievementsByIDHandler {
	return &userAchievementsByIDHandler{logger: logger, es: es}
}

type userAchievementsByIDHandler struct {
	logger logger.Logger
	es     es.AggregateStore
}

func (q *userAchievementsByIDHandler) Handle(ctx context.Context, query *GetUserAchievementsByIDQuery) ([]AchievementsProjection, error) {
	userAchievements := NewUserAchievementsAggregateWithID(query.ID)

	err := q.es.Load(ctx, userAchievements)
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	return AchievementsProjectionFromAggregate(userAchievements), nil
}

// ==

type GetUserPublicAchievementsByIDQuery struct {
	ID string
}

func NewGetUserPublicAchievementsByIDQuery(id string) *GetUserPublicAchievementsByIDQuery {
	return &GetUserPublicAchievementsByIDQuery{ID: id}
}

type UserPublicAchievementsByIDHandler interface {
	Handle(ctx context.Context, query *GetUserPublicAchievementsByIDQuery) ([]AchievementsProjection, error)
}

func NewUserPublicAchievementsByIDHandler(logger logger.Logger, es es.AggregateStore) UserPublicAchievementsByIDHandler {
	return &userPublicAchievementsByIDHandler{logger: logger, es: es}
}

type userPublicAchievementsByIDHandler struct {
	logger logger.Logger
	es     es.AggregateStore
}

func (q *userPublicAchievementsByIDHandler) Handle(ctx context.Context, query *GetUserPublicAchievementsByIDQuery) ([]AchievementsProjection, error) {
	userAchievements := NewUserAchievementsAggregateWithID(query.ID)

	err := q.es.Load(ctx, userAchievements)
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	return PublicAchievementsProjectionFromAggregate(userAchievements), nil
}
