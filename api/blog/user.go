package blog

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	model "api.seaotterms.com/model/blog"
)

type UserDataForUpdate struct {
	UpdatedAt  time.Time
	UpdateName string
	Avatar     string
}

type apiAccount struct {
	Username string
	Email    string
}

func CreateUser(c *fiber.Ctx, db *gorm.DB) error {
	var data dto.RegisterRequest
	var find []apiAccount

	if err := c.BodyParser(&data); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "客戶端資料錯誤",
		})
	}

	err := db.Model(&model.User{}).Find(&find).Error
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	data.Username = strings.ToLower(data.Username)
	data.Email = strings.ToLower(data.Email)
	// check Username & Email exist
	for _, col := range find {
		if data.Username == col.Username {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     "用戶已註冊",
			})
		} else if data.Email == col.Email {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     "電子信箱已註冊",
			})
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	data.Password = string(hashedPassword)
	dataCreate := model.User{
		Username:   data.Username,
		Password:   data.Password,
		Email:      data.Email,
		CreateName: data.Username,
	}
	err = db.Create(&dataCreate).Error
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    "註冊成功",
	})
}

func UpdateUser(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData dto.UserUpdateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     err.Error(),
		})
	}
	// URL decoding
	id, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     err.Error(),
		})
	}
	// check if form id equal route id
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}
	if u != uint64(clientData.ID) {
		logrus.Error("ID比對失敗")
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     "ID比對失敗",
		})
	}

	err = db.Model(&model.User{}).Where("id = ?", id).
		Select("updated_at", "update_name", "avatar").
		Updates(UserDataForUpdate{
			UpdatedAt:  time.Now(),
			UpdateName: clientData.Username,
			Avatar:     clientData.Avatar,
		}).Error
	if err != nil {
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			logrus.Error(err)
			return c.Status(fiber.StatusNotFound).JSON(dto.CommonResponse[any]{
				StatusCode: 404,
				ErrMsg:     "使用者不存在，更新使用者失敗",
			})
		} else {
			logrus.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
				StatusCode: 500,
				ErrMsg:     err.Error(),
			})
		}
	}
	logrus.Infof("個人資料 %s 更新成功", clientData.Username)
	return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[any]{
		StatusCode: 200,
		InfoMsg:    fmt.Sprintf("資料 %s 更新成功", clientData.Username),
	})
}
