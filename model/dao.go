package model

import "gorm.io/gorm"

type Latee struct {
	gorm.Model
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	AddedAt int64
}
