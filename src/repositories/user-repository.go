package repositories

import (
	"gin_fw/src/database"
	"gin_fw/src/models"
)

type IUserRepository interface {
	FindAll() ([]models.User, error)
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := database.DB.Model(&models.User{}).Preload("Department").Find(&users).Error
	return users, err
}
