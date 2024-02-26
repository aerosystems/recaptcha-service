package rest

type RecaptchaUsecase interface {
	VerifyV3(token string, ip string) error
}
