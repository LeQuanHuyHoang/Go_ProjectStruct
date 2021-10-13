package handler

import (
	"crawl-data/pkg/service"
)

type UserHandler struct {
	UserSrv service.IUserService
}

func NewUserHandler(srv service.IUserService) IUserHandler {
	return &UserHandler{}
}

type IUserHandler interface {
}

//Take request, call service to handle logic
