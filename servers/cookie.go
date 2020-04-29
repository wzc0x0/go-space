package servers

import (
	"github.com/gin-gonic/gin"

	"go-space/config"
)



func SetCookie(c *gin.Context,name string,value string,)  {
	conf := config.GetConfig()
	c.SetCookie(name,value, conf.Cookie.MaxAge, conf.Cookie.Path, conf.Cookie.Domain, 
		conf.Cookie.Secure,
		conf.Cookie.HttpOnly)
}

func DelCookie(name string,c *gin.Context)  {
	if cookie, err := c.Request.Cookie(name); err == nil {
		c.SetCookie(cookie.Name,cookie.Value,-1,
			cookie.Path,cookie.Domain,cookie.Secure,cookie.HttpOnly)
	}
}
