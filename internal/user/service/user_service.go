package service

import (
	"errors"
	"strconv"

	"reseller-jh-be/internal/user/model"
	"reseller-jh-be/internal/user/repository"
	"reseller-jh-be/internal/user/request"
	"reseller-jh-be/internal/user/response"
	"reseller-jh-be/pkg/common"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	Repo repository.UserRepository
}

type UserServiceInterface interface {
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: *repo,
	}
}

func (s *UserService) Login(c *gin.Context, reqUser request.ReqLogin) (resp response.RespLogin, err error) {
	var user model.User
	user, err = s.Repo.GetUserByUsername(reqUser.Username)
	if err != nil {
		return resp, errors.New("Username tidak ditemukan.")
	}

	checkPassword := common.CheckPassword(user.Password, reqUser.Password)
	if checkPassword != true {
		return resp, errors.New("Password salah.")
	}

	randomString := common.RandString(10)
	session := sessions.Default(c)
	session.Clear()
	session.Set("token", randomString)
	session.Save()
	generateToken, err := common.Encrypt(randomString)
	if err != nil {
		return resp, err
	}

	resp.ID = user.ID
	resp.Username = user.Username
	resp.Email = user.Email
	resp.Token = generateToken

	return resp, nil
}

func (s *UserService) CreateUser(reqUser request.ReqRegister) (user model.User, err error) {
	user, _ = s.Repo.GetUserByEmail(reqUser.Email)
	if user.Email != "" {
		return user, errors.New("Email sudah terdaftar.")
	}

	user, _ = s.Repo.GetUserByUsername(reqUser.Username)
	if user.Email != "" {
		return user, errors.New("Username sudah terdaftar.")
	}

	if reqUser.Password != reqUser.PasswordConfirmation {
		return user, errors.New("Password tidak cocok.")
	}

	password, err := common.HashPassword(reqUser.Password)
	if err != nil {
		return user, errors.New("Hash password error.")
	}

	reqUser.Password = password
	user, err = s.Repo.CreateUser(reqUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (user model.User, err error) {
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return user, err
	}

	user, err = s.Repo.GetUser(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
