package model

import (
	"gorm.io/gorm"
	"time"
)

type Latee struct {
	gorm.Model
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	AddedAt time.Time
}
