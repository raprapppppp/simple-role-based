package models

type Task struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	AccountId uint   `gorm:"not null"`
	Task      string `gorm:"not null"`
}