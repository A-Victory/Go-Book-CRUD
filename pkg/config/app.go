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
	d, err := gorm.Open("mysql", "admin:mypassword@tcp(database-2.c7jtdbiq3jv9.eu-central-1.rds.amazonaws.com:3306)/test1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

// GetDB just returns the database
func GetDB() *gorm.DB {
	return db
}
