package models

type WeChatSignDto struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce string `form:"nonce"`
	Echostr string `form:"echostr"`
}

type LoginDto struct {
	RedirectUrl string `form:"redirect_url" binding:"required"`
}

type LoginCode struct {
	Code string `form:"code" binding:"required" json:"code"`
	State string `form:"state" binding:"required" json:"state"`
}

type WxUser struct {
	Openid string `json:"openid"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn int `json:"expires_in"`
}

