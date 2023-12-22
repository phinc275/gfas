package gfas

func DefaultAchievements() map[AchievementID]map[AchievementTier]*Achievement {
	return map[AchievementID]map[AchievementTier]*Achievement{
		AchievementIDCompleteJobSmartContract: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteJobSmartContract,
				Category:      AchievementCategorySmartContract,
				Tier:          AchievementTierGold,
				Sort:          270100,
				BadgeName:     "Senior Blockchain Engineer",
				Description:   "Complete 15 Smart Contract jobs",
				LoyaltyPoints: 5000,
				Threshold:     15,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobSmartContract,
				Category:      AchievementCategorySmartContract,
				Tier:          AchievementTierSilver,
				Sort:          170100,
				BadgeName:     "Intermediate Smart Contract Developer",
				Description:   "Complete 5 Smart Contract jobs",
				LoyaltyPoints: 2000,
				Threshold:     5,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobSmartContract,
				Category:      AchievementCategorySmartContract,
				Tier:          AchievementTierBronze,
				Sort:          70100,
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
				Sort:          270200,
				BadgeName:     "Senior Developer",
				Description:   "Complete 25 IT Development jobs",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobITDevelopment,
				Category:      AchievementCategoryITDevelopment,
				Tier:          AchievementTierSilver,
				Sort:          170200,
				BadgeName:     "Intermediate SmartDeveloper",
				Description:   "Complete 10 IT Development jobs",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobITDevelopment,
				Category:      AchievementCategoryITDevelopment,
				Tier:          AchievementTierBronze,
				Sort:          70200,
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
				Sort:          270300,
				BadgeName:     "Senior Creative Director",
				Description:   "Complete 25 Design & Creative jobs",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobDesignAndCreative,
				Category:      AchievementCategoryDesignAndCreative,
				Tier:          AchievementTierSilver,
				Sort:          170300,
				BadgeName:     "Intermediate Graphic Designer",
				Description:   "Complete 10 Design & Creative jobs",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobDesignAndCreative,
				Category:      AchievementCategoryDesignAndCreative,
				Tier:          AchievementTierBronze,
				Sort:          70300,
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
				Sort:          270400,
				BadgeName:     "Senior Business Development Director",
				Description:   "Complete 25 Sales & MKT jobs",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobSalesAndMarketing,
				Category:      AchievementCategorySalesAndMarketing,
				Tier:          AchievementTierSilver,
				Sort:          170400,
				BadgeName:     "Intermediate Marketing Manager",
				Description:   "Complete 10 Sales & MKT jobs",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobSalesAndMarketing,
				Category:      AchievementCategorySalesAndMarketing,
				Tier:          AchievementTierBronze,
				Sort:          70400,
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
				Sort:          270500,
				BadgeName:     "Senior Blockchain Strategist",
				Description:   "Complete 25 KOL & Web3 Advisor jobs",
				LoyaltyPoints: 5000,
				Threshold:     25,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteJobKOLAndWeb3Advisor,
				Category:      AchievementCategoryKOLAndWeb3Advisor,
				Tier:          AchievementTierSilver,
				Sort:          170500,
				BadgeName:     "Intermediate Crypto Consultant",
				Description:   "Complete 10 KOL & Web3 Advisor jobs",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteJobKOLAndWeb3Advisor,
				Category:      AchievementCategoryKOLAndWeb3Advisor,
				Tier:          AchievementTierBronze,
				Sort:          70500,
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
				Sort:              250100,
				BadgeName:         "GitHub Legend",
				Description:       "Get A+ rank or higher in the GitHub stats for the Drésumé system",
				LoyaltyPoints:     5000,
				Threshold:         1,
				InternalThreshold: 3,
				Metadata:          make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:                AchievementIDStatsGithub,
				Category:          AchievementCategoryGithubStats,
				Tier:              AchievementTierSilver,
				Sort:              150100,
				BadgeName:         "GitHub Pro",
				Description:       "Get B rank or higher in the GitHub stats for the Drésumé system",
				LoyaltyPoints:     2000,
				Threshold:         1,
				InternalThreshold: 2,
				Metadata:          make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:                AchievementIDStatsGithub,
				Category:          AchievementCategoryGithubStats,
				Tier:              AchievementTierBronze,
				Sort:              50100,
				BadgeName:         "GitHub Beginner",
				Description:       "Get C rank or higher in the GitHub stats for the Drésumé system",
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
				Sort:          201400,
				BadgeName:     "Legend Freelancer",
				Description:   "Apply for 1000 jobs",
				LoyaltyPoints: 5000,
				Threshold:     1000,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDApplyJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				Sort:          101400,
				BadgeName:     "Hard-working Freelancer",
				Description:   "Apply for 100 jobs",
				LoyaltyPoints: 2000,
				Threshold:     100,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDApplyJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				Sort:          1400,
				BadgeName:     "New Freelancer",
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
				Sort:          1000,
				BadgeName:     "Hello World!",
				Description:   "First time sign-in",
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
				Sort:          100500,
				BadgeName:     "Verified Twitter User",
				Description:   "Sync twitter profile",
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
				Sort:          100300,
				BadgeName:     "Verified GitHub User",
				Description:   "Sync GitHub profile",
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
				Sort:          100400,
				BadgeName:     "Verified LinkedIn User",
				Description:   "Sync LinkedIn profile",
				LoyaltyPoints: 2000,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDVisitSite: {
			AchievementTierGold: &Achievement{
				ID:                         AchievementIDVisitSite,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierGold,
				Sort:                       201100,
				BadgeName:                  "Loyal Visitor",
				Description:                "Visit DLancer 60 days in a row",
				LoyaltyPoints:              5000,
				Threshold:                  1,
				InternalThreshold:          60,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:                         AchievementIDVisitSite,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierSilver,
				Sort:                       101100,
				BadgeName:                  "Regular Visitor",
				Description:                "Visit DLancer 20 days in a row",
				LoyaltyPoints:              2000,
				Threshold:                  1,
				InternalThreshold:          20,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:                         AchievementIDVisitSite,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierBronze,
				Sort:                       1100,
				BadgeName:                  "Hello, Visitor!",
				Description:                "Visit DLancer 3 days in a row",
				LoyaltyPoints:              500,
				Threshold:                  1,
				InternalThreshold:          3,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
		},
		AchievementIDVisitSiteTotal: {
			AchievementTierGold: &Achievement{
				ID:                         AchievementIDVisitSiteTotal,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierGold,
				Sort:                       249900,
				BadgeName:                  "Old Friend",
				Description:                "Visit DLancer 365 days in total",
				LoyaltyPoints:              100000,
				Threshold:                  1,
				InternalThreshold:          365,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
		},

		AchievementIDCompleteWorkspace: {
			AchievementTierGold: &Achievement{
				ID:            AchievementIDCompleteWorkspace,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierGold,
				Sort:          201800,
				BadgeName:     "Web3 Legend Boss",
				Description:   "Complete 50 workspaces as Client",
				LoyaltyPoints: 5000,
				Threshold:     50,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDCompleteWorkspace,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				Sort:          101800,
				BadgeName:     "Senior Web3 Recruiter",
				Description:   "Complete 20 workspaces as Client",
				LoyaltyPoints: 2000,
				Threshold:     20,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDCompleteWorkspace,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				Sort:          1800,
				BadgeName:     "Startup Entrepreneur",
				Description:   "Complete 5 workspaces as Client",
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
				Sort:          201700,
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
				Sort:          101700,
				BadgeName:     "Senior Job Leader",
				Description:   "Post 10 jobs",
				LoyaltyPoints: 2000,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDPostJob,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				Sort:          1700,
				BadgeName:     "Job Innovator",
				Description:   "Post 1 job",
				LoyaltyPoints: 500,
				Threshold:     1,
				Metadata:      make(map[string]interface{}),
			},
		},

		AchievementIDSpendMoney: {
			AchievementTierGold: &Achievement{
				ID:                         AchievementIDSpendMoney,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierGold,
				Sort:                       201600,
				BadgeName:                  "Legendary Client",
				Description:                "Spend total $50k",
				LoyaltyPoints:              5000,
				Threshold:                  1,
				InternalThreshold:          50000,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:                         AchievementIDSpendMoney,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierSilver,
				Sort:                       101600,
				BadgeName:                  "A Big Client",
				Description:                "Spend total $5k",
				LoyaltyPoints:              2000,
				Threshold:                  1,
				InternalThreshold:          5000,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:                         AchievementIDSpendMoney,
				Category:                   AchievementCategoryCommon,
				Tier:                       AchievementTierBronze,
				Sort:                       1600,
				BadgeName:                  "Hello, Contributor!",
				Description:                "Spend total $100",
				LoyaltyPoints:              500,
				Threshold:                  1,
				InternalThreshold:          100,
				ShouldShowInternalProgress: true,
				Metadata:                   make(map[string]interface{}),
			},
		},

		AchievementIDVerifyProfileIdentity: {
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDVerifyProfileIdentity,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				Sort:          1100,
				BadgeName:     "General Verified User",
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
				Sort:          1200,
				BadgeName:     "Payment Verified User",
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
				Sort:              1500,
				BadgeName:         "Loyal Grinder",
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
				Sort:          201300,
				BadgeName:     "Legendary Viewer",
				Description:   "View other profiles 100 times",
				LoyaltyPoints: 5000,
				Threshold:     100,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierSilver: &Achievement{
				ID:            AchievementIDViewOtherProfile,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierSilver,
				Sort:          101300,
				BadgeName:     "Networking Enthusiast",
				Description:   "View other profiles 50 times",
				LoyaltyPoints: 2000,
				Threshold:     50,
				Metadata:      make(map[string]interface{}),
			},
			AchievementTierBronze: &Achievement{
				ID:            AchievementIDViewOtherProfile,
				Category:      AchievementCategoryCommon,
				Tier:          AchievementTierBronze,
				Sort:          1300,
				BadgeName:     "Insightful Observer",
				Description:   "View other profiles 10 times",
				LoyaltyPoints: 500,
				Threshold:     10,
				Metadata:      make(map[string]interface{}),
			},
		},
	}
}
