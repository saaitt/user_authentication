package model

import "time"

type User struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName  string
	LastName   string
	NationalID string
	Cancelled  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

