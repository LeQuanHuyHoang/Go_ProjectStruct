package handler

import (
	"net/http"
	"project-struct/pkg/model"
	"project-struct/pkg/service"

	"gitlab.com/goxp/cloud0/ginext"
)

type AuthHandler struct {
	AuthSrv service.IAuthService
	JWTSrv  service.IJWTService
}

func NewAuthHandler(jwt service.IJWTService, auth service.IAuthService) *AuthHandler {
	return &AuthHandler{
		AuthSrv: auth,
		JWTSrv:  jwt,
	}
}

func (h *AuthHandler) Login(c *ginext.Request) (*ginext.Response, error) {
	req := model.LoginRequest{}
	c.MustBind(&req)

	user, err := h.AuthSrv.CheckUserPassword(req.Email, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusUnauthorized, err.Error())
	}
	//Get JWT
	token, err := h.JWTSrv.GenJWTToken(user.ID)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, token), nil
}

func (h *AuthHandler) SignUp(c *ginext.Request) (*ginext.Response, error) {
	rep := model.UserRequest{}
	c.MustBind(&rep)

	rs, err := h.AuthSrv.SignUp(rep.Email, rep.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
