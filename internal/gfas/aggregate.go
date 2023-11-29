package gfas

import (
	"context"

	"github.com/phinc275/gfas/pkg/es"
	"github.com/pkg/errors"
)

const (
	UserAchievementsAggregateType es.AggregateType = "USER_ACHIEVEMENTS"
)

type UserAchievementsAggregate struct {
	*es.AggregateBase
	UserAchievements *UserAchievements
}

func NewUserAchievementsAggregateWithID(id string) *UserAchievementsAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewUserAchievementsAggregate()
	aggregate.SetID(id)
	aggregate.UserAchievements.ID = id
	return aggregate
}

func NewUserAchievementsAggregate() *UserAchievementsAggregate {
	userAchievementsAggregate := &UserAchievementsAggregate{UserAchievements: NewUserAchievements()}
	base := es.NewAggregateBase(userAchievementsAggregate.When)
	base.SetType(UserAchievementsAggregateType)
	userAchievementsAggregate.AggregateBase = base
	return userAchievementsAggregate
}

func (a *UserAchievementsAggregate) When(evt es.Event) error {
	switch evt.GetEventType() {
	case EventTypeUserAchievementProgressChanged:
		return a.onUserAchievementProgressChanged(evt)
	case EventTypeUserAchievementProgressReset:
		return a.onUserAchievementProgressReset(evt)
	case EventTypeUserAchievementCompleted:
		return a.onUserAchievementCompleted(evt)
	case EventTypeUserAchievementClaimed:
		return a.onUserAchievementClaimed(evt)

	default:
		return es.ErrInvalidEventType
	}
}

func (a *UserAchievementsAggregate) onUserAchievementProgressChanged(evt es.Event) error {
	var eventData UserAchievementProgressChangedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	achievement, ok := a.UserAchievements.Achievements[eventData.AchievementID][eventData.AchievementTier]
	if !ok {
		return nil //
	}

	achievement.Progress += eventData.ProgressChanged
	achievement.InternalProgress += eventData.InternalProgressChanged
	achievement.LastObserved = evt.Timestamp

	return nil
}

func (a *UserAchievementsAggregate) onUserAchievementProgressReset(evt es.Event) error {
	var eventData UserAchievementProgressResetEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	achievement, ok := a.UserAchievements.Achievements[eventData.AchievementID][eventData.AchievementTier]
	if !ok {
		return nil //
	}

	achievement.Progress = 0
	achievement.InternalProgress = 0
	achievement.LastObserved = evt.Timestamp

	return nil
}

func (a *UserAchievementsAggregate) onUserAchievementCompleted(evt es.Event) error {
	var eventData UserAchievementCompletedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	achievement, ok := a.UserAchievements.Achievements[eventData.AchievementID][eventData.AchievementTier]
	if !ok {
		return nil //
	}

	achievement.AchievedAt = &eventData.CompletedAt

	return nil
}

func (a *UserAchievementsAggregate) onUserAchievementClaimed(evt es.Event) error {
	var eventData UserAchievementClaimedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	achievement, ok := a.UserAchievements.Achievements[eventData.AchievementID][eventData.AchievementTier]
	if !ok {
		return nil //
	}

	achievement.AchievedAt = &eventData.ClaimedAt

	return nil
}

func LoadUserAchievementsAggregate(ctx context.Context, eventStore es.AggregateStore, aggregateID string) (*UserAchievementsAggregate, error) {
	userAchievements := NewUserAchievementsAggregateWithID(aggregateID)
	err := eventStore.Exists(ctx, userAchievements.GetID())
	if err != nil {
		return userAchievements, err
	}

	err = eventStore.Load(ctx, userAchievements)
	if err != nil {
		return userAchievements, err
	}

	return userAchievements, nil
}
