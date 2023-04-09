package user

import "clean/models"

type UserUsecase interface {
	Create(user *models.User) (*models.User, error)
	ReadAll()([]*models.User, error)
	GetId(id int64) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id int64) (*models.User, error)
}