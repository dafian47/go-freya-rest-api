package config

import (
	model "github.com/dafian47/go-freya-rest-api/module/event"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
)

func InitDB(databaseUrl string, isDebug bool) *gorm.DB {

	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	db.LogMode(isDebug)

	db.AutoMigrate(&model.Event{})

	return db
}
