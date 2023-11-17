package gfas

import "time"

type ExternalEvent any

type ExternalEventJobCompleted struct {
	Timestamp   time.Time
	UserID      string
	JobId       string
	JobCategory string
}

type ExternalEventJobApplied struct {
	Timestamp   time.Time
	UserID      string
	JobId       string
	JobCategory string
}

type ExternalEventJobPosted struct {
	Timestamp   time.Time
	UserID      string
	JobID       string
	JobCategory string
}

type ExternalEventSocialConnected struct {
	Timestamp time.Time
	UserID    string
	Provider  string
}

type ExternalEventSocialRankingUpdated struct {
	Timestamp time.Time
	UserID    string
	Provider  string
	Rank      string
}

type ExternalEventUserAccessed struct {
	Timestamp time.Time
	UserID    string
}

type ExternalEventWorkspaceCompleted struct {
	Timestamp   time.Time
	UserID      string
	WorkspaceID string
}

type ExternalEventMoneySpent struct {
	Timestamp time.Time
	UserID    string
	Amount    float64
}

type ExternalEventProfileVerified struct {
	Timestamp  time.Time
	UserID     string
	VerifyType string
}

type ExternalEventProfileViewed struct {
	Timestamp time.Time
	UserID    string
}

type ExternalEventLoyaltyPointsEarned struct {
	Timestamp time.Time
	UserID    string
	Amount    int64
}
