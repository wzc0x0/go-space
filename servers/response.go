package servers

import (
	"github.com/gin-gonic/gin"

	"go-space/constants"
)

type customResponse struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg  string `json:"msg"`
}

func ResponseOK(c *gin.Context, data interface{})  {
	c.JSON(200,&customResponse{
		Code: constants.SUCCESS,
		Data: data,
		Msg: "success!",
	})
}

func ResponseException(c *gin.Context,code int,msg string)  {
	c.JSON(200,&customResponse{
		Code: code,
		Data: nil,
		Msg: msg,
	})
}
