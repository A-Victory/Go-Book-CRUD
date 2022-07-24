package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

//Connect func help us connect to our database
//
func Connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

//GetDB just returns the database
func GetDB() *gorm.DB {
	return db
}
