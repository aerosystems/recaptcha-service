package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Validate godoc
// @Summary validate reCAPTCHA v3 token
// @Tags api-gateway-special
// @Accept  json
// @Produce application/json
// @Param X-Recaptcha-V3-Token header string true "reCAPTCHA v3 token"
// @Param X-Recaptcha-V2-Token header string true "reCAPTCHA v2 token"
// @Success 204 {object} Response
// @Failure 401 {object} ErrResponse
// @Router /v1/validate [get]
func (h *BaseHandler) Validate(c echo.Context) error {
	return SuccessResponse(c, http.StatusNoContent, "", nil)
}

// ValidateV3 godoc
// @Summary validate reCAPTCHA v3 token
// @Tags api-gateway-special
// @Accept  json
// @Produce application/json
// @Param X-Recaptcha-V3-Token header string true "reCAPTCHA v3 token"
// @Success 204 {object} Response
// @Failure 401 {object} ErrResponse
// @Router /v1/validate/v3 [get]
func (h *BaseHandler) ValidateV3(c echo.Context) error {
	headers := c.Request().Header
	if err := h.recaptchaService.VerifyV3(headers.Get("X-Recaptcha-V3-Token"), c.RealIP()); err != nil {
		h.log.Info(headers.Get("X-Recaptcha-V3-Token"))
		h.log.WithError(err).Error("reCAPTCHA v3 token is not valid")
		return ErrorResponse(c, http.StatusUnauthorized, "reCAPTCHA v3 token is not valid", err)
	}

	return SuccessResponse(c, http.StatusNoContent, "", nil)
}
