package repository

import (		
	"reseller-jh-be/internal/user/model"
	"reseller-jh-be/internal/user/request"	

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (r *UserRepository) CreateUser(reqUser request.ReqRegister) (user model.User, err error) {
	user.Email = reqUser.Email
	user.Username = reqUser.Username
	user.Password = reqUser.Password	
	err = r.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetUser(ID int64) (user model.User, err error) {
	if err := r.DB.Where("id = ? ", ID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (user model.User, err error) {	
	if err := r.DB.Where("email = ? ", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (user model.User, err error) {	
	if err := r.DB.Where("username = ? ", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}