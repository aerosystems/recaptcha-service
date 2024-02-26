package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ValidateHandler struct {
	*BaseHandler
	recaptchaUsecase RecaptchaUsecase
}

func NewValidateHandler(
	baseHandler *BaseHandler,
	recaptchaUsecase RecaptchaUsecase,
) *ValidateHandler {
	return &ValidateHandler{
		BaseHandler:      baseHandler,
		recaptchaUsecase: recaptchaUsecase,
	}
}

// Validate godoc
// @Summary validate reCAPTCHA v2 token & v3 token
// @Tags api-gateway-special
// @Accept  json
// @Produce application/json
// @Param X-Recaptcha-V2-Token header string true "reCAPTCHA v2 token"
// @Param X-Recaptcha-V3-Token header string true "reCAPTCHA v3 token"
// @Success 204 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /v1/validate [get]
func (vh ValidateHandler) Validate(c echo.Context) error {
	return vh.SuccessResponse(c, http.StatusNotImplemented, "", nil)
}

// ValidateV2 godoc
// @Summary validate reCAPTCHA v2 token
// @Tags api-gateway-special
// @Accept  json
// @Produce application/json
// @Param X-Recaptcha-V2-Token header string true "reCAPTCHA v2 token"
// @Success 204 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /v2/validate [get]
func (vh ValidateHandler) ValidateV2(c echo.Context) error {
	return vh.SuccessResponse(c, http.StatusNotImplemented, "", nil)
}

// ValidateV3 godoc
// @Summary validate reCAPTCHA v3 token
// @Tags api-gateway-special
// @Accept  json
// @Produce application/json
// @Param X-Recaptcha-V3-Token header string true "reCAPTCHA v3 token"
// @Success 204 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /v3/validate [get]
func (vh ValidateHandler) ValidateV3(c echo.Context) error {
	headers := c.Request().Header
	if err := vh.recaptchaUsecase.VerifyV3(headers.Get("X-Recaptcha-V3-Token"), c.RealIP()); err != nil {
		return vh.ErrorResponse(c, http.StatusUnauthorized, "reCAPTCHA v3 token is not valid", err)
	}
	return vh.SuccessResponse(c, http.StatusNoContent, "", nil)
}
