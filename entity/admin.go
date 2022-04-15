package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        string
	Username  string `gorm:"not null"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
