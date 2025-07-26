package blog

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/galgame"
)

func QueryBrand(c *fiber.Ctx, db *gorm.DB) error {
	var responseData []model.Brand

	err := db.Order("COALESCE(updated_at, created_at) DESC").Find(&responseData).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	logrus.Info("Brand資料查詢成功")
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.Brand]{
		StatusCode: 200,
		InfoMsg:    "Brand資料查詢成功",
		Data:       &responseData,
	})
}

func CreateBrand(c *fiber.Ctx, db *gorm.DB) error {
	var requestData dto.BrandCreateRequest

	if err := c.BodyParser(&requestData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	data := model.Brand{
		Name:        requestData.Name,
		WorkAmount:  requestData.WorkAmount,
		Dissolution: requestData.Dissolution,
		CreatedAt:   time.Now(),
		CreatedName: "seaotterms",
	}

	if err := db.Create(&data).Error; err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	logrus.Info("Galgame品牌資料建立成功: " + requestData.Name)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    "Galgame品牌資料建立成功: " + requestData.Name,
	})
}
