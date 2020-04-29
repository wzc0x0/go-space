package servers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"go-space/config"
)

var DB *gorm.DB

func InitDB() (*gorm.DB,error) {
	conf := config.GetConfig()
	db,err := gorm.Open(
		conf.Database.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Name))

	if err == nil {
		db.DB().SetMaxIdleConns(conf.Database.MaxIdleConns)
		db.SingularTable(true)
		DB = db
		return db,err
	}
	return nil,err
}