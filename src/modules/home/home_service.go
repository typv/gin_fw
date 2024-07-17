package home

import (
	"gin_fw/src/models"
	"gin_fw/src/repositories"
)

type HomeService struct {
	userRepo   repositories.IUserRepository
	departRepo repositories.IDepartmentRepository
}

func NewHomeService() *HomeService {
	return &HomeService{
		userRepo:   repositories.NewUserRepository(),
		departRepo: repositories.NewDepartmentRepository(),
	}
}

func (hs *HomeService) GetHello() string {
	return "Hello from HomeController!!!"
}

func (hs *HomeService) GetUsers() []models.User {
	users, _ := hs.userRepo.FindAll()
	return users
}

func (hs *HomeService) GetDepartments() []models.Department {
	departs, _ := hs.departRepo.FindAll()
	return departs
}
