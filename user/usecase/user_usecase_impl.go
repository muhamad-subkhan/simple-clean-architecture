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


func (u *UserUsecaseImpl) GetId(id int64) (*models.User, error) {
	return u.UserRepo.GetId(id)
}


func (u *UserUsecaseImpl) Update(user *models.User) (*models.User, error) {
	return u.UserRepo.Update(user)
}

func (u *UserUsecaseImpl) Delete(id int64) (*models.User, error) {
	return u.UserRepo.Delete(id)
}