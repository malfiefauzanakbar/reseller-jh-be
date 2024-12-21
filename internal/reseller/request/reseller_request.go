package request

type ReqReseller struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Type      string `form:"type"`
}

type ReqCreateReseller struct {
	Fullname        string `form:"fullname"`
	WhatsappNo      string `form:"whatsapp_no"`
	WhatsappLink    string `form:"whatsapp_link"`
	Email           string `form:"email"`
	WhereDidYouKnow string `form:"where_did_you_know"`
	ReasonsToJoin   string `form:"reasons_to_join"`
	Address         string `form:"address"`
	CaptchaToken    string `form:"captcha_token"`
}
