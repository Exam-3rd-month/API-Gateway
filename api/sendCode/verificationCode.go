package sendcode

import (
	"crypto/rand"
	"encoding/base64"
	"gateway/internal/config"
	"log"
	"net/smtp"
)

func generateRandomCode(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

func SendVerificationCode(to string) (string, error) {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	from := configs.ResetPassword.From
	password := configs.ResetPassword.Password
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	code, err := generateRandomCode(5)
	if err != nil {
		return "", err
	}

	subject := "Subject: Verification Code\n"
	body := "Your verification code is: " + code
	msg := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return "", err
	}

	return code, nil
}
