package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

func QueryArticle(c *fiber.Ctx, db *gorm.DB) error {
	var articleData []model.Article

	result := db.Order("created_at desc").Find(&articleData)
	if result.Error != nil {
		logrus.Error(result.Error)
		// 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": articleData,
	})
}

func QuerySingleArticle(c *fiber.Ctx, db *gorm.DB) error {
	var articleData model.Article

	// find articles
	result := db.First(&articleData, c.Params("articleID"))
	if result.Error != nil {
		// if record not exist
		if result.Error == gorm.ErrRecordNotFound {
			logrus.Info(result.Error)
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": result.Error.Error(),
			})
		}
	}
	logrus.Info("查詢單一文章成功")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": articleData,
	})
}

func CreateArticle(c *fiber.Ctx, db *gorm.DB) error {
	var clientData dto.ArticleQueryRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// get tag data
	var tags []model.Tag
	if len(clientData.TagIDs) > 0 {
		if err := db.Where("id IN ?", clientData.TagIDs).Find(&tags).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
		if len(tags) != len(clientData.TagIDs) {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "部分標籤不存在",
			})
		}
	}

	data := model.Article{
		Title:   clientData.Title,
		Content: clientData.Content,
		Tags:    tags,
	}

	if err := db.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "建立資料成功",
	})
}
