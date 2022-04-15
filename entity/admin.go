package entity

import "time"

type Admin struct {
	ID        string
	Username  string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
