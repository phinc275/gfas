package gfas

import (
	"sort"
	"strings"
	"time"
)

type AchievementProjection struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Tier     string `json:"tier"`
	Sort     int64  `json:"sort"`

	BadgeName     string `json:"badge_name"`
	Description   string `json:"description"`
	LoyaltyPoints int64  `json:"loyalty_points"`

	Progress  int64 `json:"progress"`
	Threshold int64 `json:"threshold"`

	InternalProgress           float64 `json:"internal_progress,omitempty"`
	InternalThreshold          float64 `json:"internal_threshold,omitempty"`
	ShouldShowInternalProgress bool    `json:"should_show_internal_progress"`

	AchievedAt *time.Time `json:"achieved_at"`
	ClaimedAt  *time.Time `json:"claimed_at"`
}

func AchievementsProjectionFromAggregate(aggregate *UserAchievementsAggregate) []AchievementProjection {
	projections := make([]AchievementProjection, 0)

	for _, achievementTiers := range aggregate.UserAchievements.Achievements {
		for _, achievement := range achievementTiers {

			achievementProjection := AchievementProjection{
				ID:            string(achievement.ID),
				Category:      string(achievement.Category),
				Tier:          string(achievement.Tier),
				BadgeName:     achievement.BadgeName,
				Description:   achievement.Description,
				LoyaltyPoints: achievement.LoyaltyPoints,
				Progress:      achievement.Progress,
				Threshold:     achievement.Threshold,
				AchievedAt:    achievement.AchievedAt,
				ClaimedAt:     achievement.ClaimedAt,
			}

			if achievement.ShouldShowInternalProgress {
				achievementProjection.ShouldShowInternalProgress = true
				achievementProjection.InternalProgress = achievement.InternalProgress
				achievementProjection.InternalThreshold = achievement.InternalThreshold
			}

			projections = append(projections, achievementProjection)
		}
	}

	sort.Slice(projections, func(i, j int) bool {
		if projections[i].Sort != projections[j].Sort {
			return projections[i].Sort < projections[j].Sort
		}

		if projections[i].Tier != projections[j].Tier {
			m := map[string]int{AchievementTierBronze: 1, AchievementTierSilver: 2, AchievementTierGold: 3}
			iTier := m[projections[i].Tier]
			jTier := m[projections[j].Tier]
			return iTier < jTier
		}

		return strings.Compare(projections[i].ID, projections[j].ID) < 0
	})

	return projections
}

func PublicAchievementsProjectionFromAggregate(aggregate *UserAchievementsAggregate) []AchievementProjection {
	projections := make([]AchievementProjection, 0)

	for _, achievementTiers := range aggregate.UserAchievements.Achievements {
		for _, achievement := range achievementTiers {

			if !achievement.IsAchieved() {
				continue
			}

			achievementProjection := AchievementProjection{
				ID:            string(achievement.ID),
				Category:      string(achievement.Category),
				Tier:          string(achievement.Tier),
				BadgeName:     achievement.BadgeName,
				Description:   achievement.Description,
				LoyaltyPoints: achievement.LoyaltyPoints,
				Progress:      achievement.Progress,
				Threshold:     achievement.Threshold,
				AchievedAt:    achievement.AchievedAt,
				ClaimedAt:     achievement.ClaimedAt,
			}

			if achievement.ShouldShowInternalProgress {
				achievementProjection.ShouldShowInternalProgress = true
				achievementProjection.InternalProgress = achievement.InternalProgress
				achievementProjection.InternalThreshold = achievement.InternalThreshold
			}

			projections = append(projections, achievementProjection)
		}
	}

	sort.Slice(projections, func(i, j int) bool {
		if projections[i].Sort != projections[j].Sort {
			return projections[i].Sort < projections[j].Sort
		}

		if projections[i].Tier != projections[j].Tier {
			m := map[string]int{AchievementTierBronze: 1, AchievementTierSilver: 2, AchievementTierGold: 3}
			iTier := m[projections[i].Tier]
			jTier := m[projections[j].Tier]
			return iTier < jTier
		}

		return strings.Compare(projections[i].ID, projections[j].ID) < 0
	})

	return projections
}
