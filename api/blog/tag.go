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
	var responseData []model.Tag

	err := db.Order("created_at desc").Find(&responseData).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     err.Error(),
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}

	logrus.Info("Tag資料查詢成功")
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.Tag]{
		StatusCode: 200,
		InfoMsg:    "Tag資料查詢成功",
		Data:       &responseData,
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
