package gfas

import "time"

type (
	AchievementID       string
	AchievementTier     string
	AchievementCategory string
)

const (
	AchievementCategorySmartContract     AchievementCategory = "SMART_CONTRACT"
	AchievementCategoryITDevelopment     AchievementCategory = "IT_DEVELOPMENT"
	AchievementCategoryDesignAndCreative AchievementCategory = "DESIGN_AND_CREATIVE"
	AchievementCategorySalesAndMarketing AchievementCategory = "SALES_AND_MARKETING"
	AchievementCategoryKOLAndWeb3Advisor AchievementCategory = "KOL_AND_WEB3_ADVISOR"
	AchievementCategoryGithubStats       AchievementCategory = "GITHUB_STATS"
	AchievementCategoryTwitterStats      AchievementCategory = "GITHUB_TWITTER"
	AchievementCategoryLinkedInStats     AchievementCategory = "GITHUB_LINKEDIN"
	AchievementCategoryCommon            AchievementCategory = "COMMON"
)

const (
	AchievementIDCompleteJobSmartContract     AchievementID = "COMPLETE_JOB_SMART_CONTRACT"
	AchievementIDCompleteJobITDevelopment     AchievementID = "COMPLETE_JOB_IT_DEVELOPMENT"
	AchievementIDCompleteJobDesignAndCreative AchievementID = "COMPLETE_JOB_DESIGN_AND_CREATIVE"
	AchievementIDCompleteJobSalesAndMarketing AchievementID = "COMPLETE_JOB_SALES_AND_MARKETING"
	AchievementIDCompleteJobKOLAndWeb3Advisor AchievementID = "COMPLETE_JOB_KOL_AND_WEB3_ADVISOR"

	AchievementIDStatsGithub   AchievementID = "STATS_GITHUB"
	AchievementIDStatsTwitter  AchievementID = "STATS_TWITTER"
	AchievementIDStatsLinkedIn AchievementID = "STATS_LINKEDIN"

	AchievementIDApplyJob        AchievementID = "APPLY_JOB"
	AchievementIDFirstTimeSignIn AchievementID = "FIRST_TIME_SIGN_IN"

	AchievementIDConnectSocialTwitter  AchievementID = "CONNECT_SOCIAL_TWITTER"
	AchievementIDConnectSocialLinkedin AchievementID = "CONNECT_SOCIAL_LINKEDIN"
	AchievementIDConnectSocialGithub   AchievementID = "CONNECT_SOCIAL_GITHUB"

	AchievementIDVisitSite         AchievementID = "VISIT_SITE"
	AchievementIDVisitSiteTotal    AchievementID = "VISIT_SITE_TOTAL"
	AchievementIDCompleteWorkspace AchievementID = "COMPLETE_WORKSPACE"

	AchievementIDPostJob    AchievementID = "POST_JOB"
	AchievementIDSpendMoney AchievementID = "SPEND_MONEY"

	AchievementIDVerifyProfileIdentity AchievementID = "VERIFY_PROFILE_IDENTITY"
	AchievementIDVerifyProfilePayment  AchievementID = "VERIFY_PROFILE_PAYMENT"

	AchievementIDEarnLoyaltyInOneDay AchievementID = "EARN_LOYALTY_IN_ONE_DAY"

	AchievementIDViewOtherProfile AchievementID = "VIEW_OTHER_PROFILE"
)

const (
	AchievementTierBronze = "BRONZE"
	AchievementTierSilver = "SILVER"
	AchievementTierGold   = "GOLD"
)

type UserAchievements struct {
	ID           string
	Achievements map[AchievementID]map[AchievementTier]*Achievement
}

func (ua *UserAchievements) IsNewlyAchieved(achievementID AchievementID, tier AchievementTier) bool {
	a, ok := ua.Achievements[achievementID][tier]
	if !ok {
		return false
	}
	return a.IsNewlyAchieved()
}

func (ua *UserAchievements) CheckInternal(achievementID AchievementID, tier AchievementTier) bool {
	a, ok := ua.Achievements[achievementID][tier]
	if !ok {
		return false
	}
	return a.CheckInternal()
}

type Achievement struct {
	ID       AchievementID
	Category AchievementCategory
	Tier     AchievementTier
	Sort     int64

	ShouldShowInternalProgress bool
	BadgeName                  string
	Description                string
	LoyaltyPoints              int64
	Threshold                  int64
	InternalThreshold          float64 // I believe there must be a `smarter` way to implement this but...

	AchievedAt *time.Time
	ClaimedAt  *time.Time

	Progress         int64
	InternalProgress float64
	LastObserved     time.Time
	Metadata         map[string]interface{}
}

func (a *Achievement) IsAchieved() bool {
	return a.AchievedAt != nil && !a.AchievedAt.IsZero()
}

func (a *Achievement) IsNewlyAchieved() bool {
	if a.AchievedAt != nil {
		return false
	}

	return a.Progress >= a.Threshold
}

func (a *Achievement) CheckInternal() bool {
	return a.InternalProgress >= a.InternalThreshold
}

func NewUserAchievements() *UserAchievements {
	return &UserAchievements{
		ID:           "",
		Achievements: DefaultAchievements(),
	}
}
