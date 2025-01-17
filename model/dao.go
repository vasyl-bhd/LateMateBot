package model

import (
	"gorm.io/gorm"
	"time"
)

type Latee struct {
	gorm.Model
	Name           string
	AddedAt        time.Time
	IsDeleted      bool
	DeletionReason string
}
