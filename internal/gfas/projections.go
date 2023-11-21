package gfas

import "time"

type AchievementsProjection struct {
	Category     string                  `json:"category"`
	Achievements []AchievementProjection `json:"achievements"`
}

type AchievementProjection struct {
	ID    string                      `json:"id"`
	Tiers []AchievementTierProjection `json:"tiers"`
}

type AchievementTierProjection struct {
	Tier          string `json:"tier"`
	Once          bool   `json:"once"` // if once, please display as 0/1
	BadgeName     string `json:"badge_name"`
	Description   string `json:"description"`
	LoyaltyPoints int64  `json:"loyalty_points"`
	Threshold     int64  `json:"threshold"`

	Progress   int64      `json:"progress"`
	AchievedAt *time.Time `json:"achieved_at"`
	ClaimedAt  *time.Time `json:"claimed_at"`
}

func AchievementsProjectionFromAggregate(aggregate *UserAchievementsAggregate) []AchievementsProjection {
	m := make(map[AchievementCategory]map[AchievementID][]AchievementTierProjection)
	for id, tiers := range aggregate.UserAchievements.Achievements {
		definedTiers := []AchievementTier{AchievementTierBronze, AchievementTierSilver, AchievementTierGold}
		for _, definedTier := range definedTiers {
			achievement, ok := tiers[definedTier]
			if !ok {
				continue
			}
			_, ok = m[achievement.Category]
			if !ok {
				m[achievement.Category] = make(map[AchievementID][]AchievementTierProjection)
			}

			m[achievement.Category][id] = append(m[achievement.Category][id], AchievementTierProjection{
				Tier:          string(definedTier),
				Once:          achievement.Once,
				BadgeName:     achievement.BadgeName,
				Description:   achievement.Description,
				LoyaltyPoints: achievement.LoyaltyPoints,
				Threshold:     achievement.Threshold,
				Progress:      achievement.Progress,
				AchievedAt:    achievement.AchievedAt,
				ClaimedAt:     achievement.ClaimedAt,
			})
		}
	}

	projections := make([]AchievementsProjection, 0, len(m))
	for k, v := range m {
		achievements := make([]AchievementProjection, 0, len(v))
		for ak, av := range v {
			achievements = append(achievements, AchievementProjection{ID: string(ak), Tiers: av})
		}
		projections = append(projections, AchievementsProjection{
			Category:     string(k),
			Achievements: achievements,
		})
	}

	return projections
}
