package model

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Completed   bool
	DueDate     *time.Time
	CreatedAt   time.Time
}
