package user

import "clean/models"

type UserUsecase interface {
	Create(user *models.User) (*models.User, error)
	ReadAll()([]*models.User, error)
}