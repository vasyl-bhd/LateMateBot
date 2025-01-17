package dao

import (
	"LateMateBot/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("latees.db"), &gorm.Config{}) // Assign to the global `DB`z
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&model.Latee{})

	if err != nil {
		panic("Failed to make migration")
	}

	return DB
}
