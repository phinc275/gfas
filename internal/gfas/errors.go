package gfas

import "github.com/pkg/errors"

var (
	ErrAchievementNotFound       = errors.New("achievement not found")
	ErrAchievementNotAchievedYet = errors.New("achievement not completed yet")
	ErrAchievementAlreadyClaimed = errors.New("achievement already claimed")
	ErrInvalidUserAchievement    = errors.New("user achievement not found or not completed yet")
)
