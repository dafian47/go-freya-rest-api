package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"

	eventModel "github.com/dafian47/go-freya-rest-api/module/event"
	userModel "github.com/dafian47/go-freya-rest-api/module/user"
)

func InitDB(databaseUrl string, isDebug bool) *gorm.DB {

	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	db.LogMode(isDebug)

	db.AutoMigrate(&userModel.User{}, &eventModel.Event{})

	return db
}
