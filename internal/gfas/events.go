package gfas

import (
	"time"

	"github.com/phinc275/gfas/pkg/es"
)

const (
	EventTypeUserAchievementProgressChanged = "USER_ACHIEVEMENT_PROGRESS_CHANGED"
	EventTypeUserAchievementProgressReset   = "USER_ACHIEVEMENT_PROGRESS_RESET"
	EventTypeUserAchievementCompleted       = "USER_ACHIEVEMENT_COMPLETED"
	EventTypeUserAchievementClaimed         = "USER_ACHIEVEMENT_CLAIMED"
)

type UserAchievementProgressChangedEvent struct {
	AchievementID           AchievementID          `json:"achievement_id"`
	AchievementTier         AchievementTier        `json:"achievement_tier"`
	ProgressChanged         int64                  `json:"progress_changed"`
	InternalProgressChanged float64                `json:"internal_progress_changed"`
	Timestamp               time.Time              `json:"timestamp"`
	Metadata                map[string]interface{} `json:"metadata"`
}

type UserAchievementProgressResetEvent struct {
	AchievementID   AchievementID          `json:"achievement_id,omitempty"`
	AchievementTier AchievementTier        `json:"achievement_tier,omitempty"`
	Timestamp       time.Time              `json:"timestamp"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

type UserAchievementCompletedEvent struct {
	AchievementID   AchievementID   `json:"achievement_id,omitempty"`
	AchievementTier AchievementTier `json:"achievement_tier,omitempty"`
	CompletedAt     time.Time       `json:"completed_at"`
}

type UserAchievementClaimedEvent struct {
	AchievementID   AchievementID   `json:"achievement_id,omitempty"`
	AchievementTier AchievementTier `json:"achievement_tier,omitempty"`
	ClaimedAt       time.Time       `json:"claimed_at"`
}

func NewUserAchievementProgressChangedEvent(aggregate es.Aggregate, achievementID AchievementID, tier AchievementTier, progressChanged int64, internalProgressChanged float64, ts time.Time, metadata map[string]interface{}) (es.Event, error) {
	eventData := UserAchievementProgressChangedEvent{
		AchievementID:           achievementID,
		AchievementTier:         tier,
		ProgressChanged:         progressChanged,
		InternalProgressChanged: internalProgressChanged,
		Timestamp:               ts,
		Metadata:                metadata,
	}
	event := es.NewBaseEvent(aggregate, EventTypeUserAchievementProgressChanged)
	err := event.SetJsonData(&eventData)
	if err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewUserAchievementProgressResetEvent(aggregate es.Aggregate, achievementID AchievementID, tier AchievementTier, ts time.Time, metadata map[string]interface{}) (es.Event, error) {
	eventData := UserAchievementProgressChangedEvent{
		AchievementID:   achievementID,
		AchievementTier: tier,
		Timestamp:       ts,
		Metadata:        metadata,
	}
	event := es.NewBaseEvent(aggregate, EventTypeUserAchievementProgressReset)
	err := event.SetJsonData(&eventData)
	if err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewUserAchievementCompletedEvent(aggregate es.Aggregate, achievementID AchievementID, tier AchievementTier, ts time.Time) (es.Event, error) {
	eventData := UserAchievementCompletedEvent{
		AchievementID:   achievementID,
		AchievementTier: tier,
		CompletedAt:     ts,
	}
	event := es.NewBaseEvent(aggregate, EventTypeUserAchievementCompleted)
	err := event.SetJsonData(&eventData)
	if err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewUserAchievementClaimedEvent(aggregate es.Aggregate, achievementID AchievementID, tier AchievementTier, ts time.Time) (es.Event, error) {
	eventData := UserAchievementClaimedEvent{
		AchievementID:   achievementID,
		AchievementTier: tier,
		ClaimedAt:       ts,
	}
	event := es.NewBaseEvent(aggregate, EventTypeUserAchievementClaimed)
	err := event.SetJsonData(&eventData)
	if err != nil {
		return es.Event{}, err
	}
	return event, nil
}
