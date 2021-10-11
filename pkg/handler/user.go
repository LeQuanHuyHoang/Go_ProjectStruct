package handler

import (
	"crawl-data/pkg/model"
	"crawl-data/pkg/service"
	"net/http"

	"gitlab.com/goxp/cloud0/ginext"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		Service: srv,
	}
}

//Take request, call service to handle logic

func (h *Handler) SignUp(c *ginext.Request) (*ginext.Response, error) {
	rep := model.User{}

	c.MustBind(&rep)

	rs, err := h.Service.SignUp(rep.Email, rep.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
