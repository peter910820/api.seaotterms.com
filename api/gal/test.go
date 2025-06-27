package gal

import (
	"crypto/tls"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func TestMailSend(c *fiber.Ctx) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_SEND_NOREPLY"))
	m.SetHeader("To", os.Getenv("TEST_MAIL_SEND_USER"))
	m.SetHeader("Subject", "測試信件")
	m.SetBody("text/html", `
	<h1>Hello!</h1>
	<p>這是一封 <b>HTML</b> 格式的測試信。</p>
	<a href="https://example.com">點我</a>
	`)

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("MAIL_SEND_NOREPLY"), os.Getenv("MAIL_SEND_NOREPLY_PASSWORD"))
	d.TLSConfig = &tls.Config{
		ServerName: os.Getenv("MAIL_HOST"),
		MinVersion: tls.VersionTLS12,
	}

	if err := d.DialAndSend(m); err != nil {
		logrus.Fatal("寄信失敗:", err)
	}
	logrus.Println("寄信成功")

	return c.Status(fiber.StatusOK).JSON(`{statusCode: 200}`)
}
