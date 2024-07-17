package repositories

import (
	"gin_fw/src/database"
	"gin_fw/src/models"
	"gorm.io/gorm"
)

type IDepartmentRepository interface {
	FindAll() ([]models.Department, error)
}

type DepartmentRepository struct{}

func NewDepartmentRepository() IDepartmentRepository {
	return &DepartmentRepository{}
}

func (r *DepartmentRepository) FindAll() ([]models.Department, error) {
	var departments []models.Department
	err := database.DB.Model(&models.Department{}).Preload("Users", func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("Departments")
	}).Find(&departments).Error
	return departments, err
}
