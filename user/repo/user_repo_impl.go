package repo

import (
	"clean/models"
	"clean/user"
	"fmt"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func UserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}


func (r *UserRepoImpl) Create(user *models.User) (*models.User, error) {
	err := r.DB.Save(user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepoImpl) ReadAll() ([]*models.User, error) {
	var user []*models.User

	err := r.DB.Find(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, err
	}

	return user, nil
}


func (r *UserRepoImpl) GetId(id int64) (*models.User, error) {
	var user *models.User

	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		fmt.Printf("[UserRepoImpl.GetId] error execute query %v \n", err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImpl) Update(user *models.User) (*models.User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Update] error execute query %v \n", err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImpl) Delete(id int64) (*models.User, error) {
	var user = models.User{}

	err := r.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Delete] error execute query %v \n", err)
		return nil, err
	}
	return &user, err
}