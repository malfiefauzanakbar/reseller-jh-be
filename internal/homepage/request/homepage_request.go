package request

type ReqHomepage struct {	
	BannerTitle         string         `json:"banner_title" gorm:"size:255;not null"`
	// BannerImage         string         `json:"banner_image" gorm:"size:255;not null"`
	ShortDescription    string         `json:"short_description" gorm:"type:text;not null"`
	JourneyTitle        string         `json:"journey_title" gorm:"size:255;not null"`
	JourneyDescription  string         `json:"journey_description" gorm:"type:text;not null"`
	StoreTitle          string         `json:"store_title" gorm:"size:255;not null"`
	StoreDescription    string         `json:"store_description" gorm:"type:text;not null"`
	TierTitle           string         `json:"tier_title" gorm:"size:255;not null"`
	TierDescription     string         `json:"tier_description" gorm:"type:text;not null"`
	FacilityTitle       string         `json:"facility_title" gorm:"size:255;not null"`
	FacilityDescription string         `json:"facility_description" gorm:"type:text;not null"`
	VideoTitle          string         `json:"video_title" gorm:"size:255;not null"`
	VideoDescription    string         `json:"video_description" gorm:"type:text;not null"`
	VideoLink           string         `json:"video_link" gorm:"type:text;not null"`	
}