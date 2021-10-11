package service

import (
	"crawl-data/pkg/model"
	"crawl-data/pkg/repo"
	"crawl-data/pkg/utils"
	"fmt"
)

type Service struct {
	Repo *repo.Repo
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Repo: repo,
	}
}

func (h *Service) SignUp(email string, password string) (*model.User, error) {
	//Logic base on diagram
	user, err := h.Repo.CheckEmail(email)
	//err != nil, not ErrNotFound
	if err != nil && !utils.IsErrNotFound(err) {
		//Case: db error
		return nil, err
	}
	if utils.IsErrNotFound(err) {
		//ok to create
	}
	if user != nil {
		//user existed
		//return err
		return nil, fmt.Errorf("email existed")
	}

	//create user

	newUser := &model.User{
		Email: email,
		//Ma hoa pass trc
		Password: password,
	}

	user, err = h.Repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}
