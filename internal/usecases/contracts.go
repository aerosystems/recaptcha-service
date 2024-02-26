package usecases

type RecaptchaRepository interface {
	ValidateV3(response, ip string) error
}
