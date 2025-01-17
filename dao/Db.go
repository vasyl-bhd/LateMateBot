package dao

import (
	"LateMateBot/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init() *gorm.DB {
	os.Mkdir("./db", os.ModePerm)
	var err error
	DB, err = gorm.Open(sqlite.Open("db/latees.db"), &gorm.Config{}) // Assign to the global `DB`z
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&model.Latee{})

	if err != nil {
		panic("Failed to make migration")
	}

	return DB
}
