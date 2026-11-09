package model

import "time"

type CompletedTask struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	StudentID uint
	Student   User `gorm:"foreignKey:StudentID"`
	TaskID    uint
	Task      Task `gorm:"foreignKey:TaskID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
