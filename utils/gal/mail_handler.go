package gal

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func GenerateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendRegisterEmail(to string, registerKey string) {
	registerUrl := fmt.Sprintf("%sapi/gal/register/%s/%s", os.Getenv("SITE_DOMAIN"), strings.Split(to, "@")[0], registerKey)

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_SEND_NOREPLY_ACCOUNT"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "註冊驗證信件")
	m.SetBody("text/html", fmt.Sprintf(`
			<h1>Hello!</h1>
			<p>這是一封 <b>HTML</b> 格式的註冊驗證信件。</p>
			<a href="%s">點我</a>
	`, registerUrl))

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("MAIL_SEND_NOREPLY_ACCOUNT"), os.Getenv("MAIL_SEND_NOREPLY_PASSWORD"))
	d.TLSConfig = &tls.Config{
		ServerName: os.Getenv("MAIL_HOST"),
		MinVersion: tls.VersionTLS12,
	}

	if err := d.DialAndSend(m); err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("寄信成功! 目標: %s", to)
	}
}
