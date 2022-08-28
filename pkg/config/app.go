package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect func help us connect to our database
// The SQL username is connected to an amazon web(aws) instance for mysql

func Connect() {
	d, err := gorm.Open("mysql", "Master:Iamthebest!@tcp(ec2-18-185-125-16.eu-central-1.compute.amazonaws.com:3306)/test03?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

// GetDB just returns the database
func GetDB() *gorm.DB {
	return db
}
