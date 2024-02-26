package google

type Api struct {
	recaptchaV2SecretKey string
	recaptchaV3SecretKey string
}

func NewApi(
	recaptchaV2SecretKey string,
	recaptchaV3SecretKey string,
) *Api {
	return &Api{
		recaptchaV2SecretKey: recaptchaV2SecretKey,
		recaptchaV3SecretKey: recaptchaV3SecretKey,
	}
}
