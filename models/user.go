package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Gender 	 int `form:"gender"`
}
