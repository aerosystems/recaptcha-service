package google

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"
	score         = 0.5
)

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func (a *Api) ValidateV3(response, ip string) error {
	return validate(a.recaptchaV3SecretKey, response, ip)
}

func validate(secretKey, response, ip string) error {
	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("secret", secretKey)
	q.Add("response", response)
	q.Add("remoteip", ip)
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

	if body.Score < score {
		return errors.New("lower received score than expected")
	}

	return nil
}
