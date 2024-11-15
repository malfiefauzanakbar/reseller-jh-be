package request

type ReqReseller struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Type      string `form:"type"`
}
