package blog

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

func QuerySystemTodo(c *fiber.Ctx, db *gorm.DB) error {
	// get query param
	id := c.Query("id")
	systemName := c.Query("system_name")

	var data []model.SystemTodo
	var err error
	if id == "" && systemName == "" {
		err = db.Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
	} else {
		if id != "" {
			err = db.Where("id = ?", id).Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
		} else {
			err = db.Where("system_name = ?", systemName).Order("COALESCE(updated_at, created_at) DESC").Find(&data).Error
		}
	}
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到SystemTodo資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}
	logrus.Info("查詢SystemTodo資料成功")
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[[]model.SystemTodo]{
		StatusCode: 200,
		InfoMsg:    "查詢SystemTodo資料成功: ",
		Data:       &data,
	})
}

func CreateSystemTodo(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData dto.SystemTodoCreateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	data := model.SystemTodo{
		SystemName:  clientData.SystemName,
		Title:       clientData.Title,
		Detail:      clientData.Detail,
		Status:      clientData.Status,
		Deadline:    clientData.Deadline,
		Urgency:     clientData.Urgency,
		CreatedName: clientData.CreatedName,
	}

	err := db.Create(&data).Error
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	logrus.Infof("系統代辦資料 %s 創建成功", clientData.Title)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("系統代辦資料 %s 創建成功", clientData.Title),
	})
}

func UpdateSystemTodo(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData dto.SystemTodoUpdateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	clientData.UpdatedAt = time.Now()
	err := db.Model(&model.SystemTodo{}).Where("id = ?", c.Params("id")).
		Select("system_name", "title", "detail", "status", "deadline", "urgency", "updated_at", "updated_name").
		Updates(clientData).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到該SystemTodo資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}
	logrus.Infof("SystemTodo %s 更新成功", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("SystemTodo %s 更新成功", c.Params("id")),
	})
}

func QuickUpdateSystemTodo(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData dto.QuickSystemTodoUpdateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	clientData.UpdatedAt = time.Now()
	err := db.Model(&model.SystemTodo{}).Where("id = ?", c.Params("id")).
		Select("status", "updated_at", "updated_name").
		Updates(clientData).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到該SystemTodo資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}
	logrus.Infof("SystemTodo %s 更新成功", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("SystemTodo %s 更新成功", c.Params("id")),
	})
}

func DeleteSystemTodo(c *fiber.Ctx, db *gorm.DB) error {
	err := db.Where("id = ?", c.Params("id")).Delete(&model.SystemTodo{}).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "找不到該SystemTodo資料",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}
	logrus.Infof("SystemTodo %s 刪除成功", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("SystemTodo %s 刪除成功", c.Params("id")),
	})
}
