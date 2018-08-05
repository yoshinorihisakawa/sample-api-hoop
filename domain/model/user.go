package model

import (
	"time"
)

type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	FirstName   string     `json:"first_name" validate:"required"`
	LastName    string     `json:"last_name" validate:"required"`
	MailAddress string     `json:"mail_address" validate:"required,email"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}
