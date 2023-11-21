package gfas

import (
	"context"
	"time"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/phinc275/gfas/pkg/es"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/pkg/errors"
)

type UserAchievementsCommands struct {
	applyExternalEvent   ApplyExternalEventCommandHandler
	claimUserAchievement ClaimUserAchievementCommandHandler
}

func NewUserAchievementsCommands(
	applyExternalEvent ApplyExternalEventCommandHandler,
	claimUserAchievement ClaimUserAchievementCommandHandler,
) *UserAchievementsCommands {
	return &UserAchievementsCommands{
		applyExternalEvent:   applyExternalEvent,
		claimUserAchievement: claimUserAchievement,
	}
}

type ApplyExternalEventCommandHandler interface {
	Handle(ctx context.Context, command *ApplyExternalEventCommand) error
}

type ApplyExternalEventCommand struct {
	es.BaseCommand
	event ExternalEvent
}

func NewApplyExternalEventCommand(aggregateID string, event ExternalEvent) *ApplyExternalEventCommand {
	return &ApplyExternalEventCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		event:       event,
	}
}

func NewApplyExternalEventCommandHandler(logger logger.Logger, es es.AggregateStore) ApplyExternalEventCommandHandler {
	return &applyExternalHandler{
		logger: logger,
		es:     es,
	}
}

type applyExternalHandler struct {
	logger logger.Logger
	es     es.AggregateStore
}

func (h *applyExternalHandler) Handle(ctx context.Context, command *ApplyExternalEventCommand) error {
	timer := time.NewTimer(0)
	attempt := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			err := func() error {
				userAchievements, err := LoadUserAchievementsAggregate(ctx, h.es, command.GetAggregateID())
				if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
					return err
				}

				err = userAchievements.ApplyExternalEvent(ctx, command.event)
				if err != nil {
					return err
				}

				return h.es.PessimisticSave(ctx, userAchievements)
			}()

			if err == nil {
				return nil
			}

			if !errors.Is(err, esdb.ErrWrongExpectedStreamRevision) {
				return err
			}

			delay := (1 << attempt) * 100 * time.Millisecond
			attempt++
			timer.Reset(delay)
		}
	}
}

type ClaimUserAchievementCommandHandler interface {
	Handle(ctx context.Context, command *ClaimUserAchievementCommand) error
}

type ClaimUserAchievementCommand struct {
	es.BaseCommand
	achievementID   AchievementID
	achievementTier AchievementTier
}

func NewClaimUserAchievementCommand(aggregateID string, achievementID AchievementID, achievementTier AchievementTier) *ClaimUserAchievementCommand {
	return &ClaimUserAchievementCommand{
		BaseCommand:     es.NewBaseCommand(aggregateID),
		achievementID:   achievementID,
		achievementTier: achievementTier,
	}
}

func NewClaimUserAchievementCommandHandler(logger logger.Logger, es es.AggregateStore) ClaimUserAchievementCommandHandler {
	return &claimUserAchievementHandler{
		logger: logger,
		es:     es,
	}
}

type claimUserAchievementHandler struct {
	logger logger.Logger
	es     es.AggregateStore
}

func (h *claimUserAchievementHandler) Handle(ctx context.Context, command *ClaimUserAchievementCommand) error {
	timer := time.NewTimer(0)
	attempt := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			err := func() error {
				userAchievements, err := LoadUserAchievementsAggregate(ctx, h.es, command.GetAggregateID())
				if err != nil {
					return err
				}

				err = userAchievements.Claim(ctx, command.achievementID, command.achievementTier)
				if err != nil {
					return err
				}
				return h.es.PessimisticSave(ctx, userAchievements)
			}()

			if err == nil {
				return nil
			}

			if !errors.Is(err, esdb.ErrWrongExpectedStreamRevision) {
				return err
			}

			delay := (1 << attempt) * 100 * time.Millisecond
			attempt++
			timer.Reset(delay)
		}
	}
}
