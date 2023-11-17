package gfas

import "github.com/pkg/errors"

var (
	ErrAchievementNotFound       = errors.New("achievement not found")
	ErrAchievementNotAchievedYet = errors.New("achievement not completed yet")
)
