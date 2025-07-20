package blog

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

// query article data (all or use id to query single article data)
func QueryArticle(c *fiber.Ctx, db *gorm.DB) error {
	var data []model.Article
	var err error
	response := dto.CreateDefalutCommonResponse[[]model.Article]()

	id := c.Query("id")
	if id != "" {
		err = db.Preload("Tags").First(&data, id).Error
	} else {
		err = db.Preload("Tags").Order("created_at desc").Find(&data).Error
	}
	if err != nil {
		logrus.Error(err)
		response.StatusCode = 500
		response.ErrMsg = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	logrus.Info("Article資料查詢成功")
	response.Data = &data
	response.InfoMsg = "Article資料查詢成功"
	return c.Status(fiber.StatusOK).JSON(response)
}

// create article data
func CreateArticle(c *fiber.Ctx, db *gorm.DB) error {
	var clientData dto.ArticleCreateRequest
	response := dto.CreateDefalutCommonResponse[model.Article]()

	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		response.StatusCode = 500
		response.ErrMsg = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if len(clientData.Tags) > 0 {
		var count int64
		db.Model(&model.Tag{}).Where("name IN ?", clientData.Tags).Count(&count)
		if count != int64(len(clientData.Tags)) {
			logrus.Error("缺少tags，請先建立tags")
			response.StatusCode = 500
			response.ErrMsg = "缺少tags，請先建立tags"
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	data := model.Article{
		Title:   clientData.Title,
		Content: clientData.Content,
		Tags:    []model.Tag{},
	}
	for _, tag := range clientData.Tags {
		if !(strings.TrimSpace(tag) == "") {
			data.Tags = append(data.Tags, model.Tag{Name: tag})
		}
	}

	if err := db.Create(&data).Error; err != nil {
		logrus.Error(err)
		response.StatusCode = 500
		response.ErrMsg = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	logrus.Info("Article資料建立成功: " + clientData.Title)
	response.Data = &data
	response.InfoMsg = "Article資料建立成功: " + clientData.Title
	return c.Status(fiber.StatusOK).JSON(response)
}

// Query Article data use tag name
func QueryArticleForTag(c *fiber.Ctx, db *gorm.DB) error {
	var data []model.Article
	response := dto.CreateDefalutCommonResponse[[]model.Article]()

	// URL decoding
	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		logrus.Error(err)
		response.StatusCode = 400
		response.ErrMsg = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	err = db.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("JOIN tags ON tags.name = article_tags.tag_name").
		Where("tags.name = ?", name).
		Find(&data).Error
	if err != nil {
		logrus.Error(err)
		response.StatusCode = 500
		response.ErrMsg = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	logrus.Info("查詢指定Tag的Article成功: " + name)
	response.Data = &data
	response.InfoMsg = "查詢指定Tag的Article成功" + name
	return c.Status(fiber.StatusOK).JSON(response)
}

// Delete article data
func DeleteArticle(c *fiber.Ctx, db *gorm.DB) error {
	response := dto.CreateDefalutCommonResponse[[]model.Article]()

	// URL decoding
	id, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Error(err)
		response.StatusCode = 400
		response.ErrMsg = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var article model.Article
	db.Preload("Tags").First(&article, id)

	db.Model(&article).Association("Tags").Clear()

	db.Delete(&article)

	logrus.Info("刪除Article成功" + id)
	response.InfoMsg = "刪除Article成功" + id
	return c.Status(fiber.StatusOK).JSON(response)
}
