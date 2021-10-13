package handler

import (
	"crawl-data/pkg/model"
	"crawl-data/pkg/service"
	"net/http"

	"gitlab.com/goxp/cloud0/ginext"
)

type AuthHandler struct {
	UserSrv IUser
	AuthSrv service.IAuthService
}

func NewAuthHandler(user service.IUserService, auth service.IAuthService) *AuthHandler {
	return &AuthHandler{
		UserSrv: user,
		AuthSrv: auth,
	}
}

type IUser interface {
	CheckUserPassword(email string, password string) (*model.User, error)
}

func (h *AuthHandler) Login(c *ginext.Request) (*ginext.Response, error) {
	req := model.LoginRequest{}
	c.MustBind(&req)

	user, err := h.UserSrv.CheckUserPassword(req.Email, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusUnauthorized, err.Error())
	}
	//Get JWT
	token, err := h.AuthSrv.GenJWTToken(user.ID)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, token), nil
}
