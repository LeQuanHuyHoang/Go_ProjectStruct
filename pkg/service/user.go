package service

import (
	"crawl-data/pkg/model"
	"crawl-data/pkg/repo"
	"crawl-data/pkg/utils"
	"fmt"
)

type UserService struct {
	Repo repo.IRepo
}

func NewUserService(repo repo.IRepo) IUserService {
	return &UserService{
		Repo: repo,
	}
}

type IUserService interface {
	SignUp(email string, password string) (*model.User, error)
	CheckUserPassword(email string, password string) (*model.User, error)
}

func (s *UserService) SignUp(email string, password string) (*model.User, error) {
	//Logic base on diagram
	user, err := s.Repo.CheckEmail(email)
	//err != nil, not ErrNotFound
	if err != nil && !utils.IsErrNotFound(err) {
		//Case: db error
		return nil, err
	}
	if user != nil {
		//user existed
		//return err
		return nil, fmt.Errorf("email existed")
	}

	//create user

	hashpassword := password + utils.SystemHashKey
	newUser := &model.User{
		Email: email,
		//Ma hoa pass trc
		Password: hashpassword,
	}

	user, err = s.Repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) CheckUserPassword(email string, password string) (*model.User, error) {
	//get user by email
	user, err := s.Repo.CheckEmail(email)
	if err != nil {
		return nil, err
	}

	//check request password == user.password
	//Ma hoa pass
	hashpassword := password + utils.SystemHashKey
	if hashpassword != user.Password {
		return nil, fmt.Errorf("wrong password")
	}
	//pass correct
	//return JWT token
	return user, nil
}
