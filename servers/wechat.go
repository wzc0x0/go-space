package servers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"go-space/config"
	"go-space/constants"
	"go-space/models"
)

func SaveAccessToken(code string) (*models.WxUser, error) {
	u := url.Values{}
	m := map[string]string{
		"appid":      config.GetConfig().WeChat.AppId,
		"secret":     config.GetConfig().WeChat.AppSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	for k, v := range m {
		u.Add(k, v)
	}
	var b bytes.Buffer
	b.WriteString(constants.TokenURL)
	b.WriteString(u.Encode())
	resp, _ := http.Get(b.String())
	defer resp.Body.Close()

	var token models.WxUser
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, DB.Create(&token).Error
}
