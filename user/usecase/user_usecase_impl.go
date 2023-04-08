package usecase

import (
	"clean/models"
	"clean/user"
)

type UserUsecaseImpl struct {
	UserRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase{
	return &UserUsecaseImpl{userRepo}
}


func (u *UserUsecaseImpl) Create(user *models.User) (*models.User, error) {
	return u.UserRepo.Create(user)
}

func (u *UserUsecaseImpl) ReadAll() ([]*models.User, error){
	return u.UserRepo.ReadAll()
} 