package request

type ReqReseller struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Type      string `form:"type"`
}

type ReqCreateReseller struct {	
	Fullname     string         `form:"fullname"`
	WhatsappNo   string         `form:"whatsapp_no"`
	WhatsappLink string         `form:"whatsapp_link"`
	Email        string         `form:"email"`
	NIK          string         `form:"nik"`
	Address      string         `form:"address"`	
}