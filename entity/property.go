package entity

import (
	"time"

	"gorm.io/gorm"
)

type Property struct {
	ID          string `gorm:"type:char(9)"`
	Name        string `gorm:"not null;size:80"`
	Description string `gorm:"not null"`
	Amount      uint16 `gorm:"not null"`
	GroupID     string `gorm:"type:char(5)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
