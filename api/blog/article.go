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
	var responseData []model.Article
	var err error

	articleID, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	if articleID != "" {
		err = db.Preload("Tags").First(&responseData, articleID).Error
	} else {
		err = db.Preload("Tags").Order("created_at desc").Find(&responseData).Error
	}
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	logrus.Info("Article資料查詢成功")
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.Article]{
		StatusCode: 200,
		InfoMsg:    "Article資料查詢成功",
		Data:       &responseData,
	})
}

// create article data
func CreateArticle(c *fiber.Ctx, db *gorm.DB) error {
	var clientData dto.ArticleCreateRequest

	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	if len(clientData.Tags) > 0 {
		var count int64
		db.Model(&model.Tag{}).Where("name IN ?", clientData.Tags).Count(&count)
		if count != int64(len(clientData.Tags)) {
			logrus.Error("缺少tags，請先建立tags")
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     "缺少tags，請先建立tags",
			})
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
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	logrus.Info("Article資料建立成功: " + clientData.Title)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    "Article資料建立成功: " + clientData.Title,
	})
}

// func UpdateArticle(c *fiber.Ctx, db *gorm.DB) error {
// 	// load client data
// 	var clientData
// 	if err := c.BodyParser(&clientData); err != nil {
// 		logrus.Error(err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"msg": err.Error(),
// 		})
// 	}
// 	response := dto.CreateDefalutCommonResponse[[]model.Article]()

// }

// Delete article data
func DeleteArticle(c *fiber.Ctx, db *gorm.DB) error {

	// URL decoding
	id, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     err.Error(),
		})
	}

	var article model.Article
	db.Preload("Tags").First(&article, id)

	db.Model(&article).Association("Tags").Clear()

	db.Delete(&article)

	logrus.Info("刪除Article成功" + id)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    "刪除Article成功" + id,
	})
}

// Query Article data use tag name
func QueryArticleForTag(c *fiber.Ctx, db *gorm.DB) error {
	var responseData []model.Article

	// URL decoding
	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     err.Error(),
		})
	}
	err = db.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("JOIN tags ON tags.name = article_tags.tag_name").
		Where("tags.name = ?", name).
		Find(&responseData).Error
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	logrus.Info("查詢指定Tag的Article成功: " + name)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.Article]{
		StatusCode: 200,
		InfoMsg:    "查詢指定Tag的Article成功" + name,
		Data:       &responseData,
	})
}
