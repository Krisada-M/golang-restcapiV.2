package services

import "api/models"

type UserService interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User, *string) error
	DeleteUser(*string) error
	GetOnlyUser(*string) (*models.User, error)
	GetAllUser() ([]*models.User, error)
}
