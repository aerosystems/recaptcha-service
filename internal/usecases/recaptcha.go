package usecases

type RecaptchaUsecase struct {
	recaptchaRepo RecaptchaRepository
}

func NewRecaptchaUsecase(
	recaptchaRepo RecaptchaRepository,
) *RecaptchaUsecase {
	return &RecaptchaUsecase{
		recaptchaRepo: recaptchaRepo,
	}
}
func (ru RecaptchaUsecase) VerifyV3(response, ip string) error {
	return ru.recaptchaRepo.ValidateV3(response, ip)
}
