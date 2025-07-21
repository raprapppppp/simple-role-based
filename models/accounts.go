package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Firstname string         `gorm:"not null"`
	Lastname  string         `gorm:"not null"`
	Username  string         `gorm:"unique;not null"` // must be unique
	Password  string         `gorm:"not null"`        // hashed password
	Role      string         `gorm:"type:text;default:'user'';not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // optional soft delete
}

type LoginCred struct {
	Username string
	Password string
}
