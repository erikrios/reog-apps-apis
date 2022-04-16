package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        string `gorm:"type:char(4)"`
	Username  string `gorm:"not null;size:20;unique"`
	Name      string `gorm:"not null;size:50"`
	Password  string `gorm:"not null;size:60"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
