package entity

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID         string     `gorm:"type:char(5)"`
	Name       string     `gorm:"not null;size:80"`
	Leader     string     `gorm:"not null;size:80"`
	Address    Address    `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Properties []Property `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
