package model

import "time"

type Task struct {
	ID             uint            `gorm:"primaryKey;autoIncrement"`
	Title          string          `gorm:"size:255;not null"`
	Description    string          `gorm:"type:text;not null"`
	Deadline       time.Time       `gorm:"not null"`
	CompletedTasks []CompletedTask `gorm:"foreignKey:TaskID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
