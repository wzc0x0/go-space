package routers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"go-space/constants"
	"go-space/models"
	"go-space/servers"
)

func Add(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"msg":err.Error(),
		})
		return
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(user.Password))
	cipher := md5Ctx.Sum(nil)

	c.JSON(200, gin.H{
		"name":     user.Name,
		"password": hex.EncodeToString(cipher),
	})
}

func GetAllCountries(c *gin.Context)  {
	var counties []models.Country
	var data = make(map[string]interface{})

	if err := servers.DB.Find(&counties).Error
		err != nil {
		servers.ResponseException(c,constants.ERROR,err.Error())
		return
	}
	data["list"] = counties
	servers.ResponseOK(c, data)
}

func RegisterUser(c *gin.Context)  {
	var user models.Usnidaizheer
	var userDto models.User

	if err := c.ShouldBind(&userDto); err != nil {
		servers.ResponseException(c,constants.InvalidParams,err.Error())
		return
	}

	if servers.DB.Where("name = ?",userDto.Name).First(&user).Error == nil {
		servers.ResponseException(c,1001,"name already exits")
		return
	}

	hashByte,err := bcrypt.GenerateFromPassword([]byte(userDto.Password),10)
	if err != nil {
		servers.ResponseException(c,constants.ERROR,err.Error())
		return
	}

	userDto.Password = string(hashByte)
	servers.DB.Create(&userDto)
	servers.ResponseOK(c,nil)
}

func LogonUser(c *gin.Context)  {
	var userDto models.User
	var user models.User

	if err := c.ShouldBind(&userDto); err != nil {
		servers.ResponseException(c,constants.InvalidParams,err.Error())
		return
	}

	if servers.DB.Where("name = ?",userDto.Name).Find(&user).Error != nil {
		servers.ResponseException(c,1001,"user not exist!")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(userDto.Password))
		err != nil {
		servers.ResponseException(c,1002,"user or password is not right!")
		return
	}
	servers.SetCookie(c,constants.SessionName,fmt.Sprint(user.ID))
	servers.ResponseOK(c,nil)
}

func GetUserInfo(c *gin.Context)  {
	var user models.User
	type userDto struct{
		ID uint `json:"id"`
		Name string `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Gender int `json:"gender"`
	}
	uid,err := c.Cookie(constants.SessionName)
	if err == nil && servers.DB.Where("id = ?",uid).Find(&user).Error == nil {
		servers.ResponseOK(c,&userDto{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Gender:    user.Gender,
		})
		return
	}

	servers.ResponseException(c,1001,"not found user!")
}

func Logout(c *gin.Context)  {
	servers.DelCookie(constants.SessionName,c)
	servers.ResponseOK(c,nil)
}


func md5Hash(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func bcryptHash(s string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(s),10)
	if err != nil {
		return s,err
	}
	return string(hash),nil
}