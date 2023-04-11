package infra

import (
	"github.com/SuperJourney/gopen/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Setting = config.LoadConfig()

var DB, _ = gorm.Open(sqlite.Open(Setting.DBFile), &gorm.Config{})
