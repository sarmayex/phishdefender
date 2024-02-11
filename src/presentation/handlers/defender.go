package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sarmayex/phishdefender/domain"
	"net/http"
)

type DefenderHandler struct {
	service *domain.PhishDefenderService
}

func NewDefenderHandler(service *domain.PhishDefenderService) *DefenderHandler {
	return &DefenderHandler{
		service: service,
	}
}

func (d *DefenderHandler) Store(c echo.Context) error {
	return c.JSON(http.StatusServiceUnavailable, "not implemented")
}
