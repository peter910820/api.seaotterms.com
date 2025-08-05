package blog

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	middleware "api.seaotterms.com/middleware/blog"
	model "api.seaotterms.com/model/blog"
	utils "api.seaotterms.com/utils/blog"
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
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, "客戶端資料錯誤", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := db.Model(&model.User{}).Find(&find).Error
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusOK, fmt.Sprintf("Todo %s 刪除成功", c.Params("id")), nil)
		return c.Status(fiber.StatusOK).JSON(response)
	}

	data.Username = strings.ToLower(data.Username)
	data.Email = strings.ToLower(data.Email)
	// check Username & Email exist
	for _, col := range find {
		if data.Username == col.Username {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, "用戶已註冊", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else if data.Email == col.Email {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, "電子信箱已註冊", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
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
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := utils.ResponseFactory[any](c, fiber.StatusOK, "註冊成功", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func UpdateUser(c *fiber.Ctx, db *gorm.DB, store *session.Store) error {
	// load client data
	var clientData dto.UserUpdateRequest
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	// URL decoding
	id, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	// check if form id equal route id
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if u != uint64(clientData.ID) {
		logrus.Error("ID比對失敗")
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, "ID比對失敗", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	timeNow := time.Now()
	err = db.Model(&model.User{}).Where("id = ?", id).
		Select("updated_at", "update_name", "avatar").
		Updates(UserDataForUpdate{
			UpdatedAt:  timeNow,
			UpdateName: clientData.Username,
			Avatar:     clientData.Avatar,
		}).Error
	if err != nil {
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			logrus.Error(err)
			response := utils.ResponseFactory[any](c, fiber.StatusNotFound, "使用者不存在，更新使用者失敗", nil)
			return c.Status(fiber.StatusNotFound).JSON(response)
		} else {
			logrus.Error(err)
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}
	// 更新使用者資料成功，更新快取表
	userInfo, ok := middleware.UserInfo[clientData.ID]
	if !ok {
		logrus.Error("快取表更新失敗")
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, "快取表更新失敗", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	userInfo.UpdatedAt = timeNow
	userInfo.UpdateName = clientData.Username
	userInfo.Avatar = clientData.Avatar
	userInfo.DataVersion++

	// 更新response UserInfo
	sess, err := store.Get(c)
	if err != nil {
		logrus.Fatal(err)
	}
	userID := sess.Get("id")
	if userID == nil {
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, "上下文資料更新失敗", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	c.Locals("user_info", userInfo)

	logrus.Infof("個人資料 %s 更新成功", clientData.Username)
	response := utils.ResponseFactory[any](c, fiber.StatusOK, fmt.Sprintf("資料 %s 更新成功", clientData.Username), nil)
	return c.Status(fiber.StatusOK).JSON(response)
}
