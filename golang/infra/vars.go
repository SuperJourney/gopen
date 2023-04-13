package infra

import (
	"github.com/SuperJourney/gopen/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Setting = config.LoadConfig()

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(Setting.DBFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
