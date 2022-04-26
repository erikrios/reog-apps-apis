package entity

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID           string `gorm:"type:char(5)"`
	Address      string `gorm:"not null"`
	VillageID    string `gorm:"not null;type:char(10)"`
	VillageName  string `gorm:"not null;size:255"`
	DistrictID   string `gorm:"not null;type:char(7)"`
	DistrictName string `gorm:"not null;size:255"`
	RegencyID    string `gorm:"not null;type:char(4)"`
	RegencyName  string `gorm:"not null;size:255"`
	ProvinceID   string `gorm:"not null;type:char(2)"`
	ProvinceName string `gorm:"not null;size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
