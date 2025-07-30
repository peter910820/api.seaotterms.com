package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/galgame"
)

func QueryGameRecord(c *fiber.Ctx, db *gorm.DB) error {
	var responseData []model.PlayRecord

	err := db.Order("COALESCE(updated_at, created_at) DESC").Find(&responseData).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	logrus.Info("個人攻略Galgame攻略資料查詢成功")
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.PlayRecord]{
		StatusCode: 200,
		InfoMsg:    "個人攻略Galgame攻略資料查詢成功",
		Data:       &responseData,
	})
}
