package captcha

type RecaptchaService struct {
	SecretKeyV2 string
	SecretKeyV3 string
}

func NewRecaptchaService(
	secretKeyV2 string,
	secretKeyV3 string,
) *RecaptchaService {
	return &RecaptchaService{
		SecretKeyV2: secretKeyV2,
		SecretKeyV3: secretKeyV3,
	}
}

func (rs *RecaptchaService) VerifyV2(token string, ip string) (bool, error) {
	return true, nil
}
