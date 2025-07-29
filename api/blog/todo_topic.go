package blog

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

func QueryTodoTopic(c *fiber.Ctx, db *gorm.DB) error {
	// URL decoding
	owner, err := url.QueryUnescape(c.Params("owner"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "客戶端資料錯誤",
		})
	}

	data := []model.TodoTopic{}
	err = db.Where("topic_owner = ?", owner).Order("topic_name DESC").Find(&data).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到TodoTopic資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}
	logrus.Info("查詢TodoTopic資料成功")
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.TodoTopic]{
		StatusCode: 200,
		InfoMsg:    "查詢TodoTopic資料成功",
		Data:       &data,
	})
}

func CreateTodoTopic(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData dto.TodoTopicCreateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	data := model.TodoTopic{
		TopicName:  clientData.TopicName,
		TopicOwner: clientData.TopicOwner,
		UpdateName: clientData.UpdateName,
	}

	err := db.Create(&data).Error
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	logrus.Infof("資料 %s 創建成功", clientData.TopicName)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("資料 %s 創建成功", clientData.TopicName),
	})
}
