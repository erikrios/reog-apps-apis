package entity

import (
	"time"

	"gorm.io/gorm"
)

type ShowSchedule struct {
	ID        string    `gorm:"type:char(9)"`
	GroupID   string    `gorm:"type:char(5)"`
	Place     string    `gorm:"not null"`
	StartOn   time.Time `gorm:"not null"`
	FinishOn  time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
