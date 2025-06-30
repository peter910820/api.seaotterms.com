package blog

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

type TagData struct {
	ID    uint
	Title string
}

func QueryTag(c *fiber.Ctx, db *gorm.DB) error {
	var tagData []model.Tag

	result := db.Order("id desc").Find(&tagData)
	if result.Error != nil {
		// if record not exist
		if result.Error == gorm.ErrRecordNotFound {
			logrus.Info(result.Error)
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			logrus.Fatal(result.Error)
		}
	}
	logrus.Info("查詢全部tag資料成功")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": tagData,
	})
}

func CreateTag(c *fiber.Ctx, db *gorm.DB) error {
	var clientData dto.TagCreateRequest

	// load client data
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	dataCreate := model.Tag{
		Name: clientData.Name,
	}
	if strings.TrimSpace(clientData.IconName) != "" {
		dataCreate.IconName = clientData.IconName
	}
	err := db.Create(&dataCreate).Error
	if err != nil {
		logrus.Println("錯誤:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "成功建立Tag",
	})
}
