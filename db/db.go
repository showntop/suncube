package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

func init() {
	var err error

	DB, err = gorm.Open("postgres", "host=localhost user=showntop dbname=suncube_dev sslmode=disable password=1")

	if err != nil {
		panic(err)
	}

}
