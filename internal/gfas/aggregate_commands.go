package gfas

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (a *UserAchievementsAggregate) ApplyExternalEvent(ctx context.Context, extEvent ExternalEvent) error {
	switch e := extEvent.(type) {
	case *ExternalEventJobCompleted:
		return a.applyExternalEventJobCompleted(ctx, e)
	case *ExternalEventJobApplied:
		return a.applyExternalEventJobApplied(ctx, e)
	case *ExternalEventJobPosted:
		return a.applyExternalEventJobPosted(ctx, e)
	case *ExternalEventSocialConnected:
		return a.applyExternalEventSocialConnected(ctx, e)
	case *ExternalEventSocialRankingUpdated:
		return a.applyExternalEventSocialRankingUpdated(ctx, e)
	case *ExternalEventUserAccessed:
		return a.applyExternalEventUserAccessed(ctx, e)
	case *ExternalEventWorkspaceCompleted:
		return a.applyExternalEventWorkspaceCompleted(ctx, e)
	case *ExternalEventMoneySpent:
		return a.applyExternalEventMoneySpent(ctx, e)
	case *ExternalEventProfileVerified:
		return a.applyExternalEventProfileVerified(ctx, e)
	case *ExternalEventProfileViewed:
		return a.applyExternalEventProfileViewed(ctx, e)
	case *ExternalEventLoyaltyPointsEarned:
		return a.applyExternalEventLoyaltyPointsEarned(ctx, e)

	default:
		return nil // not interested in other kind of event
	}
}

func (a *UserAchievementsAggregate) Claim(ctx context.Context, achievementID AchievementID, tier AchievementTier) error {
	achievement, ok := a.UserAchievements.Achievements[achievementID][tier]
	if !ok {
		return ErrAchievementNotFound
	}

	if !achievement.IsAchieved() {
		return ErrAchievementNotAchievedYet
	}

	if achievement.ClaimedAt != nil {
		return ErrAchievementAlreadyClaimed
	}

	now := time.Now()
	event, err := NewUserAchievementClaimedEvent(a, achievementID, tier, now)
	if err != nil {
		return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
	}
	err = a.Apply(event)
	if err != nil {
		return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
	}

	externalEvent := ExternalEventLoyaltyPointsEarned{
		Timestamp: now,
		UserID:    a.UserAchievements.ID,
		Amount:    achievement.LoyaltyPoints,
	}

	return a.ApplyExternalEvent(ctx, externalEvent)
}

func (a *UserAchievementsAggregate) applyExternalEventJobCompleted(_ context.Context, e *ExternalEventJobCompleted) error {
	var achievementID AchievementID

	switch e.JobCategory {
	case string(AchievementCategorySmartContract):
		achievementID = AchievementIDCompleteJobSmartContract
	case string(AchievementCategoryITDevelopment):
		achievementID = AchievementIDCompleteJobITDevelopment
	case string(AchievementCategoryDesignAndCreative):
		achievementID = AchievementIDCompleteJobDesignAndCreative
	case string(AchievementCategorySalesAndMarketing):
		achievementID = AchievementIDCompleteJobSalesAndMarketing
	case string(AchievementCategoryKOLAndWeb3Advisor):
		achievementID = AchievementIDCompleteJobKOLAndWeb3Advisor
	default:
		return nil // not interested in other job category, yet
	}

	return a.applySimpleProgressChanged(false, achievementID, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventJobApplied(_ context.Context, e *ExternalEventJobApplied) error {
	return a.applySimpleProgressChanged(false, AchievementIDApplyJob, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventJobPosted(_ context.Context, e *ExternalEventJobPosted) error {
	return a.applySimpleProgressChanged(false, AchievementIDPostJob, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventSocialConnected(_ context.Context, e *ExternalEventSocialConnected) error {
	var achievementID AchievementID

	switch strings.ToLower(e.Provider) {
	case "github":
		achievementID = AchievementIDConnectSocialGithub
	case "twitter":
		achievementID = AchievementIDConnectSocialTwitter
	case "linkedin":
		achievementID = AchievementIDConnectSocialLinkedin
	default:
		return nil
	}

	return a.applySimpleProgressChanged(true, achievementID, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventSocialRankingUpdated(_ context.Context, e *ExternalEventSocialRankingUpdated) error {
	var (
		achievementID    AchievementID
		internalProgress float64
	)

	switch strings.ToUpper(e.Rank) {
	case "S+", "S", "S-", "A+":
		internalProgress = 3
	case "A", "A-", "B+", "B":
		internalProgress = 2
	default:
		internalProgress = 1
	}

	switch strings.ToLower(e.Provider) {
	case "github":
		achievementID = AchievementIDStatsGithub
	case "twitter":
		achievementID = AchievementIDStatsTwitter
	case "linkedin":
		achievementID = AchievementIDStatsLinkedIn
	default:
		return nil
	}

	for _, tier := range []AchievementTier{AchievementTierBronze, AchievementTierSilver, AchievementTierGold} {
		achievement, ok := a.UserAchievements.Achievements[achievementID][tier]
		if !ok {
			continue
		}

		if achievement.IsAchieved() {
			continue
		}

		if internalProgress < achievement.InternalProgress { // this definitely is cheating
			continue
		}

		event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 0, internalProgress-achievement.InternalProgress, e.Timestamp, map[string]interface{}{})
		if err != nil {
			return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
		}
		err = a.Apply(event)
		if err != nil {
			return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
		}

		if a.UserAchievements.CheckInternal(achievementID, tier) {
			event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 1, 0, e.Timestamp, map[string]interface{}{})
			if err != nil {
				return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
			}
		}

		if a.UserAchievements.IsNewlyAchieved(achievementID, tier) {
			event, err = NewUserAchievementCompletedEvent(a, achievementID, tier, e.Timestamp)
			if err != nil {
				return errors.Wrap(err, "Apply:NewUserAchievementCompletedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementCompletedEvent")
			}
		}
	}

	return nil
}

func (a *UserAchievementsAggregate) applyExternalEventUserAccessed(_ context.Context, e *ExternalEventUserAccessed) error {
	var achievementID AchievementID

	achievementID = AchievementIDFirstTimeSignIn
	err := a.applySimpleProgressChanged(true, achievementID, 1, e.Timestamp)
	if err != nil {
		return err
	}

	achievementID = AchievementIDVisitSite
	for _, tier := range []AchievementTier{AchievementTierBronze, AchievementTierSilver, AchievementTierGold} {
		achievement, ok := a.UserAchievements.Achievements[achievementID][tier]
		if !ok {
			continue
		}

		if achievement.IsAchieved() {
			continue
		}

		// FIXME: use configurable timezone
		lastObservedDate := time.Date(
			achievement.LastObserved.UTC().Year(),
			achievement.LastObserved.UTC().Month(),
			achievement.LastObserved.UTC().Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)

		newDate := time.Date(
			achievement.LastObserved.UTC().Year(),
			achievement.LastObserved.UTC().Month(),
			achievement.LastObserved.UTC().Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)

		if lastObservedDate.After(newDate) {
			continue
		}

		if newDate.Sub(lastObservedDate) > 24*time.Hour {
			event, err := NewUserAchievementProgressResetEvent(a, achievementID, tier, e.Timestamp, map[string]interface{}{})
			if err != nil {
				return errors.Wrap(err, "NewUserAchievementProgressResetEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementProgressResetEvent")
			}
			return nil
		}

		event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 0, 1, e.Timestamp, map[string]interface{}{})
		if err != nil {
			return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
		}
		err = a.Apply(event)
		if err != nil {
			return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
		}

		if a.UserAchievements.CheckInternal(achievementID, tier) {
			event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 1, 0, e.Timestamp, map[string]interface{}{})
			if err != nil {
				return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
			}
		}

		if a.UserAchievements.IsNewlyAchieved(achievementID, tier) {
			event, err = NewUserAchievementCompletedEvent(a, achievementID, tier, e.Timestamp)
			if err != nil {
				return errors.Wrap(err, "Apply:NewUserAchievementCompletedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementCompletedEvent")
			}
		}
	}

	return nil
}

func (a *UserAchievementsAggregate) applyExternalEventWorkspaceCompleted(_ context.Context, e *ExternalEventWorkspaceCompleted) error {
	return a.applySimpleProgressChanged(false, AchievementIDCompleteWorkspace, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventMoneySpent(_ context.Context, e *ExternalEventMoneySpent) error {
	achievementID := AchievementIDSpendMoney
	for _, tier := range []AchievementTier{AchievementTierBronze, AchievementTierSilver, AchievementTierGold} {
		achievement, ok := a.UserAchievements.Achievements[achievementID][tier]
		if !ok {
			continue
		}

		if achievement.IsAchieved() {
			continue
		}

		event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 0, e.Amount, e.Timestamp, map[string]interface{}{})
		if err != nil {
			return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
		}
		err = a.Apply(event)
		if err != nil {
			return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
		}

		if a.UserAchievements.CheckInternal(achievementID, tier) {
			event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 1, 0, e.Timestamp, map[string]interface{}{})
			if err != nil {
				return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
			}
		}

		if a.UserAchievements.IsNewlyAchieved(achievementID, tier) {
			event, err = NewUserAchievementCompletedEvent(a, achievementID, tier, e.Timestamp)
			if err != nil {
				return errors.Wrap(err, "Apply:NewUserAchievementCompletedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementCompletedEvent")
			}
		}
	}

	return nil
}

func (a *UserAchievementsAggregate) applyExternalEventProfileVerified(_ context.Context, e *ExternalEventProfileVerified) error {
	var achievementID AchievementID
	switch e.VerifyType {
	case "payment":
		achievementID = AchievementIDVerifyProfilePayment
	case "identity":
		achievementID = AchievementIDVerifyProfileIdentity
	default:
		return nil
	}

	return a.applySimpleProgressChanged(false, achievementID, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventProfileViewed(_ context.Context, e *ExternalEventProfileViewed) error {
	return a.applySimpleProgressChanged(false, AchievementIDViewOtherProfile, 1, e.Timestamp)
}

func (a *UserAchievementsAggregate) applyExternalEventLoyaltyPointsEarned(_ context.Context, e *ExternalEventLoyaltyPointsEarned) error {
	achievementID := AchievementIDEarnLoyaltyInOneDay
	for _, tier := range []AchievementTier{AchievementTierBronze, AchievementTierSilver, AchievementTierGold} {
		achievement, ok := a.UserAchievements.Achievements[achievementID][tier]
		if !ok {
			continue
		}

		if achievement.IsAchieved() {
			continue
		}

		// FIXME: use configurable timezone
		lastObservedDate := time.Date(
			achievement.LastObserved.UTC().Year(),
			achievement.LastObserved.UTC().Month(),
			achievement.LastObserved.UTC().Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)

		newDate := time.Date(
			achievement.LastObserved.UTC().Year(),
			achievement.LastObserved.UTC().Month(),
			achievement.LastObserved.UTC().Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)

		if lastObservedDate.After(newDate) {
			continue
		}

		if newDate.Sub(lastObservedDate) > 24*time.Hour {
			event, err := NewUserAchievementProgressResetEvent(a, achievementID, tier, e.Timestamp, map[string]interface{}{})
			if err != nil {
				return errors.Wrap(err, "NewUserAchievementProgressResetEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementProgressResetEvent")
			}
			return nil
		}

		event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 0, 1, e.Timestamp, map[string]interface{}{})
		if err != nil {
			return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
		}
		err = a.Apply(event)
		if err != nil {
			return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
		}

		if a.UserAchievements.CheckInternal(achievementID, tier) {
			event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, 1, 0, e.Timestamp, map[string]interface{}{})
			if err != nil {
				return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
			}
		}

		if a.UserAchievements.IsNewlyAchieved(achievementID, tier) {
			event, err = NewUserAchievementCompletedEvent(a, achievementID, tier, e.Timestamp)
			if err != nil {
				return errors.Wrap(err, "Apply:NewUserAchievementCompletedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementCompletedEvent")
			}
		}
	}

	return nil
}

func (a *UserAchievementsAggregate) applySimpleProgressChanged(skipIfAchieved bool, achievementID AchievementID, progressChanged int64, ts time.Time) error {
	for _, tier := range []AchievementTier{AchievementTierBronze, AchievementTierSilver, AchievementTierGold} {
		achievement, ok := a.UserAchievements.Achievements[achievementID][tier]
		if !ok {
			continue
		}

		if skipIfAchieved && achievement.IsAchieved() {
			continue
		}

		event, err := NewUserAchievementProgressChangedEvent(a, achievementID, tier, progressChanged, 0, ts, map[string]interface{}{})
		if err != nil {
			return errors.Wrap(err, "NewUserAchievementProgressChangedEvent")
		}
		err = a.Apply(event)
		if err != nil {
			return errors.Wrap(err, "Apply:UserAchievementProgressChangedEvent")
		}

		if achievement.IsNewlyAchieved() {
			event, err = NewUserAchievementCompletedEvent(a, achievementID, tier, ts)
			if err != nil {
				return errors.Wrap(err, "Apply:NewUserAchievementCompletedEvent")
			}
			err = a.Apply(event)
			if err != nil {
				return errors.Wrap(err, "Apply:UserAchievementCompletedEvent")
			}
		}
	}

	return nil
}
