package gfas

func DefaultAchievements() map[AchievementID]map[AchievementTier]*Achievement {
	return map[AchievementID]map[AchievementTier]*Achievement{
		AchievementIDCompleteJobSmartContract: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteJobSmartContract,
				Category:      AchievementCategorySmartContract,
				Tier:          AchievementTierGold,
				BadgeName:     "Senior Blockchain Engineer",
				Description:   "Complete 15 Smart Contract job",
				LoyaltyPoints: 5000,
				Threshold:     15,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobSmartContract,
				Category:      AchievementCategorySmartContract,
				Tier:          AchievementTierSilver,
				BadgeName:     "Intermediate Smart Contract Developer",
				Description:   "Complete 5 Smart Contract job",
				LoyaltyPoints: 2000,
				Threshold:     5,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobSmartContract,
				Category:      AchievementCategorySmartContract,
				Tier:          AchievementTierBronze,
				BadgeName:     "Junior Smart Contract Developer",
				Description:   "Complete 1 Smart Contract job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDCompleteJobITDevelopment: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteJobITDevelopment,
				Category:      AchievementCategoryITDevelopment,
				Tier:          AchievementTierGold,
				BadgeName:     "Senior Developer",
				Description:   "Complete 25 IT Development job",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobITDevelopment,
				Category:      AchievementCategoryITDevelopment,
				Tier:          AchievementTierSilver,
				BadgeName:     "Intermediate SmartDeveloper",
				Description:   "Complete 10 IT Development job",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobITDevelopment,
				Category:      AchievementCategoryITDevelopment,
				Tier:          AchievementTierBronze,
				BadgeName:     "Junior SmartDeveloper",
				Description:   "Complete 1 IT Development job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDCompleteJobDesignAndCreative: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteJobDesignAndCreative,
				Category:      AchievementCategoryDesignAndCreative,
				Tier:          AchievementTierGold,
				BadgeName:     "Senior Creative Director",
				Description:   "Complete 25 Design & Creative job",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobDesignAndCreative,
				Category:      AchievementCategoryDesignAndCreative,
				Tier:          AchievementTierSilver,
				BadgeName:     "Intermediate Graphic Designer",
				Description:   "Complete 10 Design & Creative job",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobDesignAndCreative,
				Category:      AchievementCategoryDesignAndCreative,
				Tier:          AchievementTierBronze,
				BadgeName:     "Junior Graphic Designer",
				Description:   "Complete 1 Design & Creative job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDCompleteJobSalesAndMarketing: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteJobSalesAndMarketing,
				Category:      AchievementCategorySalesAndMarketing,
				Tier:          AchievementTierGold,
				BadgeName:     "Senior Business Development Director",
				Description:   "Complete 25 Sales & MKT job",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobSalesAndMarketing,
				Category:      AchievementCategorySalesAndMarketing,
				Tier:          AchievementTierSilver,
				BadgeName:     "Intermediate Marketing Manager",
				Description:   "Complete 10 Sales & MKT job",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobSalesAndMarketing,
				Category:      AchievementCategorySalesAndMarketing,
				Tier:          AchievementTierBronze,
				BadgeName:     "Junior Sales Representative",
				Description:   "Complete 1 Sales & MKT job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDCompleteJobKOLAndWeb3Advisor: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteJobKOLAndWeb3Advisor,
				Category:      AchievementCategoryKOLAndWeb3Advisor,
				Tier:          AchievementTierGold,
				BadgeName:     "Senior Blockchain Strategist",
				Description:   "Complete 25 KOL & Web3 Advisor job",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobKOLAndWeb3Advisor,
				Category:      AchievementCategoryKOLAndWeb3Advisor,
				Tier:          AchievementTierSilver,
				BadgeName:     "Intermediate Crypto Consultant",
				Description:   "Complete 10 KOL & Web3 Advisor job",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobKOLAndWeb3Advisor,
				Category:      AchievementCategoryKOLAndWeb3Advisor,
				Tier:          AchievementTierBronze,
				BadgeName:     "Junior Web3 Specialist",
				Description:   "Complete 1 KOL & Web3 Advisor job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDStatsGithub: {
			AchievementTierGold: &Achievement{
				ID:                AchievementIDStatsGithub,
				Category:          AchievementCategoryGithubStats,
				Tier:              AchievementTierGold,
				BadgeName:         "Github legend",
				Description:       "Get A+ rank or above in Github stats",
				Once:              true,
				LoyaltyPoints:     5000,
				Threshold:         1,
				InternalThreshold: 3,
				Metadata:          make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:                AchievementIDStatsGithub,
				Category:          AchievementCategoryGithubStats,
				Tier:              AchievementTierSilver,
				BadgeName:         "Github pro ",
				Description:       "Get B rank or above in Github stats",
				Once:              true,
				LoyaltyPoints:     2000,
				Threshold:         1,
				InternalThreshold: 2,
				Metadata:          make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:                AchievementIDStatsGithub,
				Category:          AchievementCategoryGithubStats,
				Tier:              AchievementTierBronze,
				BadgeName:         "Github beginner",
				Description:       "Get C rank or above in Github stats",
				Once:              true,
				LoyaltyPoints:     500,
				Threshold:         1,
				InternalThreshold: 1,
				Metadata:          make(map[string]interface{}),
			},
		},

		AchievementIDApplyJob: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDApplyJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierGold,
				BadgeName:     "Legend freelancer",
				Description:   "Apply for 1000 jobs",
				LoyaltyPoints: 5000,
				Threshold:     1000,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDApplyJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Hard-working freelander",
				Description:   "Apply for 100 jobs",
				LoyaltyPoints: 2000,
				Threshold:     100,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDApplyJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "New freelancer",
				Description:   "Apply for 1 jobs",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDFirstTimeSignIn: {
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDFirstTimeSignIn,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "Hello world!",
				Description:   "First time sign-in",
				Once:          true,
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDConnectSocialTwitter: {
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDConnectSocialTwitter,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Verified Twitter user",
				Description:   "Sync twitter profile",
				Once:          true,
				LoyaltyPoints: 2000,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDConnectSocialGithub: {
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDConnectSocialGithub,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Verified Github user",
				Description:   "Sync github profile",
				Once:          true,
				LoyaltyPoints: 2000,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDConnectSocialLinkedin: {
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDConnectSocialLinkedin,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Verified Linked-in user",
				Description:   "Sync linked-in profile",
				Once:          true,
				LoyaltyPoints: 2000,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDVisitSite: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDVisitSite,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierGold,
				BadgeName:     "Loyal visitor",
				Description:   "Visit the site 60 days in a row",
				LoyaltyPoints: 5000,
				Threshold:     60,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDVisitSite,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Regular visitor",
				Description:   "Visit the site 20 days in a row",
				LoyaltyPoints: 2000,
				Threshold:     20,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDVisitSite,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "Hello, visitor!",
				Description:   "Visit the site 3 days in a row",
				LoyaltyPoints: 500,
				Threshold:     3,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDCompleteWorkspace: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteWorkspace,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierGold,
				BadgeName:     "Web3 Legend Boss",
				Description:   "Complete 50 workspace as Client",
				LoyaltyPoints: 5000,
				Threshold:     50,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteWorkspace,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Senior Web3 Recruiter",
				Description:   "Complete 20 workspace as Client",
				LoyaltyPoints: 2000,
				Threshold:     20,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteWorkspace,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "Startup Entrepreneur",
				Description:   "Complete 5 workspace as Client",
				LoyaltyPoints: 500,
				Threshold:     5,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDPostJob: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDPostJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierGold,
				BadgeName:     "Web3 Project Director",
				Description:   "Post 30 jobs",
				LoyaltyPoints: 5000,
				Threshold:     30,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDPostJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Senior job leader",
				Description:   "Post 10 jobs",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDPostJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "Job Innovator",
				Description:   "Post 1 job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDSpendMoney: {
			AchievementTierGold: &Achievement{
				ID:                AchievementIDSpendMoney,
				Category:          AchievementCategoryCommon,
				Tier:              AchievementTierGold,
				BadgeName:         "Legendary client",
				Description:       "Spend total $50k",
				Once:              true,
				LoyaltyPoints:     5000,
				Threshold:         1,
				InternalThreshold: 50000,
				Metadata:          make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:                AchievementIDSpendMoney,
				Category:          AchievementCategoryCommon,
				Tier:              AchievementTierSilver,
				BadgeName:         "A big client",
				Description:       "Spend total $5k",
				Once:              true,
				LoyaltyPoints:     2000,
				Threshold:         1,
				InternalThreshold: 5000,
				Metadata:          make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:                AchievementIDSpendMoney,
				Category:          AchievementCategoryCommon,
				Tier:              AchievementTierBronze,
				BadgeName:         "Hello, contributor!",
				Description:       "Spend total $100",
				Once:              true,
				LoyaltyPoints:     500,
				Threshold:         1,
				InternalThreshold: 100,
				Metadata:          make(map[string]interface{}),
			},
		},

		AchievementIDVerifyProfileIdentity: {
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDVerifyProfileIdentity,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "General verified user",
				Description:   "Identify verified profile",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},
		AchievementIDVerifyProfilePayment: {
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDVerifyProfilePayment,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "Payment verified user",
				Description:   "Payment verified profile",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDEarnLoyaltyInOneDay: {
			AchievementTierBronze: &Achievement{
				ID:                AchievementIDEarnLoyaltyInOneDay,
				Category:          AchievementCategoryCommon,
				Tier:              AchievementTierBronze,
				BadgeName:         "Loyal grinder",
				Description:       "Get 100 loyalty in a single day",
				LoyaltyPoints:     500,
				Threshold:         1,
				InternalThreshold: 100,
				Metadata:          make(map[string]interface{}),
			},
		},

		AchievementIDViewOtherProfile: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDViewOtherProfile,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierGold,
				BadgeName:     "Legendary viewer",
				Description:   "View other profile 100 times",
				LoyaltyPoints: 5000,
				Threshold:     100,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDViewOtherProfile,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				BadgeName:     "Networking Enthusiast",
				Description:   "View other profile 50 times",
				LoyaltyPoints: 2000,
				Threshold:     50,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDViewOtherProfile,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				BadgeName:     "Insightful Observer",
				Description:   "View other profile 10 times",
				LoyaltyPoints: 500,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
		},
	}
}
