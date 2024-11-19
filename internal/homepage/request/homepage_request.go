package request

type ReqHomepage struct {	
	BannerTitle         string         `form:"banner_title"`
	// BannerImage         string         `form:"banner_imageull"`
	ShortDescription    string         `form:"short_description"`
	JourneyTitle        string         `form:"journey_title"`
	JourneyDescription  string         `form:"journey_description"`
	StoreTitle          string         `form:"store_title"`
	StoreDescription    string         `form:"store_description"`
	TierTitle           string         `form:"tier_title"`
	TierDescription     string         `form:"tier_description"`
	FacilityTitle       string         `form:"facility_title"`
	FacilityDescription string         `form:"facility_description"`
	VideoTitle          string         `form:"video_title"`
	VideoDescription    string         `form:"video_description"`
	VideoLink           string         `form:"video_link"`	
}