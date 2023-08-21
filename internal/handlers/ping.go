package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *BaseHandler) Ping(c echo.Context) error {
	return SuccessResponse(c, http.StatusOK, "pong", nil)
}
