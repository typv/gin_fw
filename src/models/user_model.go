package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	DepartmentID *uint  `gorm:"column:department_id"`
	Name         string `gorm:"column:name"` // Can Override NamingStrategy
	Email        string
	Password     *string
	DeletedAt    *time.Time
	Department   *Department `gorm:"foreignKey:DepartmentID"`
}

func (User) TableName() string {
	return "users"
}
