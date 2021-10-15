package handler

import (
	"net/http"
	"project-struct/pkg/model"
	"project-struct/pkg/service"
	"strings"

	"gitlab.com/goxp/cloud0/ginext"
)

type UserHandler struct {
	UserSrv service.IUserService
	JWTSrv  service.IJWTService
}

func NewUserHandler(srv service.IUserService, jwt service.IJWTService) IUserHandler {
	return &UserHandler{
		UserSrv: srv,
		JWTSrv:  jwt,
	}
}

type IUserHandler interface {
	Profile(c *ginext.Request) (*ginext.Response, error)
	Update(c *ginext.Request) (*ginext.Response, error)
	Delete(c *ginext.Request) (*ginext.Response, error)
}

//Take request, call service to handle logic

func (h *UserHandler) Profile(c *ginext.Request) (*ginext.Response, error) {
	authHeader := c.GinCtx.Request.Header.Get("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := h.JWTSrv.ValidateToken(authHeader)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	rs, err := h.UserSrv.Profile(userID)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *UserHandler) Update(c *ginext.Request) (*ginext.Response, error) {
	rep := model.UserUpdate{}
	c.MustBind(&rep)
	authHeader := c.GinCtx.Request.Header.Get("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := h.JWTSrv.ValidateToken(authHeader)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	rs, err := h.UserSrv.Update(userID, rep.NewPassword)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, rs.Password), nil
}

func (h *UserHandler) Delete(c *ginext.Request) (*ginext.Response, error) {
	authHeader := c.GinCtx.Request.Header.Get("Authorization")
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := h.JWTSrv.ValidateToken(authHeader)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if h.UserSrv.Delete(userID) != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, nil), nil
}
