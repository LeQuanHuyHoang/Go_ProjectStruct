package handler

import (
	"crawl-data/pkg/model"
	"crawl-data/pkg/service"
	"net/http"

	"gitlab.com/goxp/cloud0/ginext"
)

type UserHandler struct {
	UserSrv service.IUserService
}

func NewUserHandler(srv service.IUserService) IUserHandler {
	return &UserHandler{
		UserSrv: srv,
	}
}

type IUserHandler interface {
	SignUp(c *ginext.Request) (*ginext.Response, error)
}

//Take request, call service to handle logic

func (h *UserHandler) SignUp(c *ginext.Request) (*ginext.Response, error) {
	rep := model.UserRequest{}

	c.MustBind(&rep)

	rs, err := h.UserSrv.SignUp(rep.Email, rep.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
