package gal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/gal"
	"api.seaotterms.com/utils"
)

func Login(c *fiber.Ctx, db *gorm.DB) error {
	var clientData dto.LoginRequest
	var responseData dto.LoginResponse
	// load client data
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(responseData)
	}
	// check user agent
	if !utils.CheckUserAgent(c) {
		return c.Status(fiber.StatusForbidden).JSON(responseData)
	}
	// 檢查ip 決定要不要發信警告
	// ip := c.IP()

	return c.Status(fiber.StatusOK).JSON(responseData)
}

func Register(c *fiber.Ctx, db *gorm.DB) error {
	return nil
}
