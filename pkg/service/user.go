package service

import (
	"project-struct/pkg/model"
	"project-struct/pkg/repo"
	"project-struct/pkg/utils"
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
	Update(userID string, newpass string) (*model.User, error)
	Delete(userID string) error
}

func (s *UserService) Profile(userID string) (*model.User, error) {
	userProfile, err := s.Repo.ViewProfile(userID)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}

func (s *UserService) Update(userID string, newpass string) (*model.User, error) {
	hashpass, _ := utils.HashPassword(newpass)
	userupdate, err := s.Repo.UpdateUser(userID, hashpass)
	if err != nil {
		return nil, err
	}
	return userupdate, nil
}

func (s *UserService) Delete(userID string) error {
	err := s.Repo.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
