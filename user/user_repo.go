package user

import "clean/models"

type UserRepo interface {
	Create(user *models.User) (*models.User, error)
	ReadAll() ([]*models.User, error)
}