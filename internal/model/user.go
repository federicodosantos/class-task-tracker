package model

import "time"

type User struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"size:100;not null"`
	Email          string `gorm:"size:100;not null;unique"`
	Password       string `gorm:"size:255;not null"`
	RoleID         uint
	Role           Role            `gorm:"foreignKey:RoleID"`
	CompletedTasks []CompletedTask `gorm:"foreignKey:StudentID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
