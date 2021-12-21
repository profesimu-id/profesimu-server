package entity

import (
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	UserID       int `gorm:"primaryKey"`
	ProfessionID int `gorm:"primaryKey"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
