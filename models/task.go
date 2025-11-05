package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Status      string         `json:"status"`
	DueDate     *time.Time     `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
