package service

import (
	"crawl-data/pkg/model"
	"crawl-data/pkg/repo"
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
	Profile(userID string) (*model.User, error)
}

func (s *UserService) Profile(userID string) (*model.User, error) {
	userProfile, err := s.Repo.ViewProfile(userID)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}
