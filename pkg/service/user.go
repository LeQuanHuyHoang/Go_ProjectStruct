package service

import (
	"crawl-data/pkg/repo"
)

type UserService struct {
	Repo repo.IRepo
}

func NewUserService(repo repo.IRepo) IUserService {
	return &UserService{}
}

type IUserService interface {
}
