package captcha

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type RecaptchaService struct {
	SecretKeyV2 string
	SecretKeyV3 string
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
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
func (rs *RecaptchaService) VerifyV3(response, remoteIP string) error {
	return checkRecaptcha(rs.SecretKeyV3, response, remoteIP)
}

func (rs *RecaptchaService) VerifyV2(response, remoteIP string) error {
	return checkRecaptcha(rs.SecretKeyV2, response, remoteIP)
}

func checkRecaptcha(secretKey, response, remoteIP string) error {
	const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"
	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("secret", secretKey)
	q.Add("response", response)
	q.Add("remoteip", remoteIP)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var body SiteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	if !body.Success {
		return errors.New("unsuccessful recaptcha verify request")
	}

	if body.Action != "login" {
		return errors.New("mismatched recaptcha action")
	}

	if body.Score < 0.5 {
		return errors.New("lower received score than expected")
	}

	return nil
}
