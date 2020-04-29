package routers

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"

	"go-space/config"
	"go-space/constants"
	"go-space/models"
	"go-space/servers"
)

func WeChatTestSign(c *gin.Context) {
	var weChatSignDto models.WeChatSignDto
	if err := c.ShouldBind(&weChatSignDto); err != nil {
		servers.ResponseException(c, 1001, "")
		return
	}
	v := []string{constants.TokenSign, weChatSignDto.Timestamp, weChatSignDto.Nonce}
	sort.Strings(v)
	s := strings.Join(v, "")
	// middleware
	if fmt.Sprintf("%x", sha1.Sum([]byte(s))) == weChatSignDto.Signature {
		c.String(200, "%s", weChatSignDto.Echostr)
	}
}

func WeChatLogin(c *gin.Context) {
	var loginDto models.LoginDto
	if err := c.ShouldBind(&loginDto); err != nil {
		servers.ResponseException(c, constants.InvalidParams, err.Error())
		return
	}
	u := getAuthURL(loginDto.RedirectUrl)
	c.Redirect(constants.Redirect, u)
	c.GetRawData()
}

func WeChatGetToken(c *gin.Context) {
	var loginCode models.LoginCode
	if err := c.ShouldBind(&loginCode); err != nil {
		servers.ResponseException(c, constants.InvalidParams, err.Error())
		return
	}
	token, err := servers.SaveAccessToken(loginCode.Code)
	if err != nil {
		servers.ResponseException(c, constants.ERROR, err.Error())
		return
	}
	servers.SetCookie(c, constants.SessionName, token.Openid)
	c.Redirect(constants.Redirect, loginCode.State)
}

func getAuthURL(state string) string {
	u := url.Values{}
	m := map[string]string{
		"appid":         config.GetConfig().WeChat.AppId,
		"redirect_uri":  config.GetConfig().WeChat.RedirectURL,
		"response_type": "code",
		"scope":         constants.Scope,
		"state":         state,
	}
	for k, v := range m {
		u.Add(k, v)
	}
	var b bytes.Buffer
	b.WriteString(constants.AuthURL)
	b.WriteString(u.Encode())
	b.WriteString("#wechat_redirect")
	return b.String()
}
