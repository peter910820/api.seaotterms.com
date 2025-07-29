package blog

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

func QueryTodoByOwner(c *fiber.Ctx, db *gorm.DB) error {
	// URL decoding
	owner, err := url.QueryUnescape(c.Params("owner"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "客戶端資料錯誤",
		})
	}

	var responseData []model.Todo
	err = db.Where("owner = ?", owner).Order("created_at DESC").Find(&responseData).Error
	if err != nil {
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			logrus.Error(err)
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
	logrus.Infof("查詢%s的Todo資料成功", owner)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.Todo]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("查詢%s的Todo資料成功", owner),
		Data:       &responseData,
	})
}

func CreateTodo(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData model.Todo
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "客戶端資料錯誤",
		})
	}
	// handle topic
	lastSlashIndex := strings.LastIndex(clientData.Topic, "/")
	if lastSlashIndex != -1 {
		clientData.Topic = clientData.Topic[:lastSlashIndex]
	} else {
		logrus.Error("topic value has error")
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "客戶端資料轉換錯誤",
		})
	}

	data := model.Todo{
		ID:         clientData.ID,
		Owner:      clientData.Owner,
		Topic:      clientData.Topic,
		Title:      clientData.Title,
		Status:     clientData.Status,
		Deadline:   clientData.Deadline,
		CreateName: clientData.CreateName,
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
	logrus.Infof("資料 %s 創建成功", clientData.Title)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("資料 %s 創建成功", clientData.Title),
	})
}

func UpdateTodoStatus(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData dto.TodoUpdateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "客戶端資料錯誤",
		})
	}
	clientData.UpdatedAt = time.Now()
	err := db.Model(&model.Todo{}).Where("id = ?", c.Params("id")).
		Select("status", "updated_at", "update_name").
		Updates(clientData).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到該Todo資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}

	logrus.Infof("Todo %s 更新成功", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("Todo %s 更新成功", c.Params("id")),
	})
}

func DeleteTodo(c *fiber.Ctx, db *gorm.DB) error {
	err := db.Where("id = ?", c.Params("id")).Delete(&model.Todo{}).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到該Todo資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}

	logrus.Infof("Todo %s 刪除成功", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("Todo %s 刪除成功", c.Params("id")),
	})
}
