package constants

const (
	SystemName     = "example-api"
	PromSystem     = "example_api"
	NameSpace      = "example"
	TrackIDHeader  = "track_id"
	SystemUser     = "ExampleSystem"
	Undefined      = "undefined"
	QCBudgetPrefix = "QC_TEST"
)

const (
	HCMRegionID    = 1
	HNRegionID     = 2
	DAKLAKRegionID = 18 // Test Only
	DNRegionID     = 17
)

const (
	HCMRegionName    = "Hồ Chí Minh"
	HNRegionName     = "Hà Nội"
	DAKLAKRegionName = "Đăk Lăk"
	DNRegionName     = "Đà Nẵng"
)

var ListRegionIDAvailable = []int{
	HCMRegionID,
	HNRegionID,
	DAKLAKRegionID,
	DNRegionID,
}

type ctxRequestIDKey int

const (
	RequestIDKey ctxRequestIDKey = 0
)

// Những biến chung
const (
	ContextTimeout             = 3 // seconds
	TimeConfig                 = 2
	DefaultLimit               = 20
	BonusRate                  = 10
	DecimalDot                 = "."
	DeeplinkPathNumber         = 3
	WarnTime                   = 3000 // ms
	LimitCharacterCampaignCode = 50
	LimitCharacterCampaignName = 50
	LimitCharacterOperation    = 200
	LimitAbortPickingTimes     = 3 // times -> tương đương là 3 ngày
	HourInDay                  = 24
	LastHourInDay              = 23
	// TimeStampOndDay            = 10 * 60 // For TEST 10 min,
	TimeStampOndDay         = 24 * 60 * 60        // 24 hours * 60 minutes * 60 seconds
	TimeStampRevertPackage  = 2 * TimeStampOndDay // 2 * 24 hours * 60 minutes * 60 seconds
	HaflItemNumber          = 2
	MaxPercentInProgressBar = 85
	PartOfTime              = 10
	LimitDayUseFarmWallet   = 7
)

// Loại cấu hình
const (
	ConfigType = 1 // Lì xì tích lũy
)

// Loại phần thưởng
const (
	RewardType = 1 // Hoàn tiền ví Farm
)

// Trạng thái status cấu hình
const (
	ConfigStatusActive   = 1
	ConfigStatusDeactive = 2
)

// Giá trị ngẫu nhiên của tiền thưởng
const (
	MaxPercent      = 100
	MaxAwardPercent = 90
	MinAwardPercent = 60
)

// Link hình ảnh
const (
	PacketSnapBarImage            = "https://media3.scdn.vn/img4/2023/01_12/Y1L7WkqePc29IjJ1Wp0n.png"
	PacketSnapBarImageClaim       = "https://media3.scdn.vn/img4/2023/01_18/3xWHaCBXaBBi8fyiHX4Q.png"
	PacketCardInfoBackgroundImage = "https://media3.scdn.vn/img4/2023/01_09/AFTkTMGjDUkvFL7sTbCm.png"
	BottomMessageImage            = "https://media3.scdn.vn/img4/2023/01_10/uPPY4BVXKijHEtvRIKDW.png"
	LandingPageStep1Image         = "https://media3.scdn.vn/img4/2023/01_10/4TbceC7eTb0FwYHienNy.png"
	LandingPageStep1ImageSelected = "https://media3.scdn.vn/img4/2023/01_09/JhESK1vSb9iT9UrbqBCK.png"
	LandingPageStep2Image         = "https://media3.scdn.vn/img4/2023/01_10/n3GwjJbojk7g4WXBS03L.png"
	LandingPageStep2ImageSelected = "https://media3.scdn.vn/img4/2023/01_09/eq5Q5hXhtjtvrdYIBejP.png"
	LandingPageStep3Image         = "https://media3.scdn.vn/img4/2023/01_10/n3GwjJbojk7g4WXBS03L.png"
	LandingPageStep3ImageSelected = "https://media3.scdn.vn/img4/2023/01_09/eq5Q5hXhtjtvrdYIBejP.png"
	LandingPageStep4Image         = "https://media3.scdn.vn/img4/2023/01_09/THjjhKBmFHElkJq9Z0f3.png"
	LandingPageStep4ImageSelected = "https://media3.scdn.vn/img4/2023/01_10/9LFsEG2otZeNYay9diIv.png"
	LandingPageImageSuccess       = "https://media3.scdn.vn/img4/2023/01_17/tyM6nld4OJOUVN52Hbpn.png"
	BgFlareImage                  = "https://media3.scdn.vn/img4/2023/01_11/eVsfUCc3GqgJ16b1mMkJ.png"
	ButtonOpenImage               = "https://media3.scdn.vn/img4/2023/01_11/aSFMZGyxqoXDPLXeHquP.png"
	ButtonBottomImage             = "https://media3.scdn.vn/img4/2023/02_01/8RactC3U8Z1xJfN8iHxt.png"
	BoxBottomImage                = "https://media3.scdn.vn/img4/2023/01_31/0hOVBE6lV8ggEAstGilX.png"
	BottomClaimImage              = "https://media3.scdn.vn/img4/2023/01_18/psIz0G7QFgp0xD7QnoO8.png"
	BoxTopImage                   = "https://media3.scdn.vn/img4/2023/01_11/GAHZZfADi08nTiFmBTaw.png"
	BoxBackImage                  = "https://media3.scdn.vn/img4/2023/02_02/4GcSOrTtBomzrGW47jEK.png"
	ProgressIndicatorImage        = "https://media3.scdn.vn/img4/2023/01_12/Y1L7WkqePc29IjJ1Wp0n.png"
	LandingPageBgRulesCoinImage   = "https://media3.scdn.vn/img4/2023/01_12/T5fFV5HjtrWECV9k4rmD.png"
	LandingPageBgRulesImage       = "https://media3.scdn.vn/img4/2023/01_12/kVrtP4vemd9TycN5gTjY.png"
	PickPackageOverQuantityImage  = "https://media3.scdn.vn/img4/2023/01_18/TcVyPF2q5YZZioaK7Chj.png"
	IconCoinImage                 = "https://media3.scdn.vn/img4/2023/01_31/BahqTGAnVSshRQBqdvW7.png"
)

// deeplink
const (
	BottomMessageStartDeeplink  = "https://www.sendo.vn/sendofarm/li-xi"
	BottomMessageInGameDeeplink = "https://www.sendo.vn/sendofarm/chi-tiet-li-xi"
	FamrWalletDeeplink          = "https://www.sendo.vn/sendofarm/vi-farm"
)

const (
	TopupOrderCode = "OG_LuckyMoney_%s_%d"
)

// #Color
var (
	BackgroundSetPackageColor    = []string{"#FDF5DE", "#FBE7B7"}
	BottomMessageBackGroundColor = "#ffde5e4"
)

// Type campaign
const (
	SponsoredProductsType = "Sponsored Products"
)
