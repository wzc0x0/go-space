package middlewares

import (
	"github.com/gin-gonic/gin"

	"go-space/constants"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie(constants.SessionName); err == nil {
			c.SetCookie(cookie.Name,cookie.Value, cookie.MaxAge,
				cookie.Path,cookie.Domain,cookie.Secure,cookie.HttpOnly)
			c.Next()
		} else  {
			c.Abort()
			c.JSON(403,gin.H{
				"msg":"forbidden",
			})
		}
	}
}