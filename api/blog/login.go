package blog

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dto "api.seaotterms.com/dto/blog"
	middleware "api.seaotterms.com/middleware/blog"
	model "api.seaotterms.com/model/blog"
)

func Login(c *fiber.Ctx, store *session.Store, db *gorm.DB) error {
	var data dto.LoginRequest
	var databaseData []dto.LoginRequest

	if err := c.BodyParser(&data); err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CommonResponse[any]{
			StatusCode: 400,
			ErrMsg:     err.Error(),
		})
	}

	err := db.Model(&model.User{}).Find(&databaseData).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
			StatusCode: 500,
			ErrMsg:     err.Error(),
		})
	}

	data.Username = strings.ToLower(data.Username)
	for _, col := range databaseData {
		// 找到使用者
		if data.Username == col.Username {
			err := bcrypt.CompareHashAndPassword([]byte(col.Password), []byte(data.Password))
			if err != nil {
				if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
					logrus.Error("login error: password not correct")
					return c.Status(fiber.StatusUnauthorized).JSON(dto.CommonResponse[any]{
						StatusCode: 401,
						ErrMsg:     "密碼輸入錯誤",
					})
				} else {
					logrus.Error(err)
					return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
						StatusCode: 500,
						ErrMsg:     err.Error(),
					})
				}
			}
			// set session
			sess, err := store.Get(c)
			if err != nil {
				logrus.Fatal(err) // 這邊之後會發送訊息
			}
			sess.Set("username", data.Username)
			if err := sess.Save(); err != nil {
				logrus.Fatal(err) // 這邊之後會發送訊息
			}
			logrus.Infof("Username %s login success", data.Username)

			var userData model.User

			err = db.Where("username = ?", data.Username).First(&userData).Error
			if err != nil {
				logrus.Error(err)
				return c.Status(fiber.StatusInternalServerError).JSON(dto.CommonResponse[any]{
					StatusCode: 500,
					ErrMsg:     err.Error(),
				})
			}

			data := middleware.UserData{
				ID:         userData.ID,
				Username:   userData.Username,
				Email:      userData.Email,
				Exp:        userData.Exp,
				Management: userData.Management,
				CreatedAt:  userData.CreatedAt,
				UpdatedAt:  userData.UpdatedAt,
				UpdateName: userData.UpdateName,
				Avatar:     userData.Avatar,
			}

			return c.Status(fiber.StatusOK).JSON(dto.CommonResponse[middleware.UserData]{
				StatusCode: 200,
				InfoMsg:    fmt.Sprintf("SystemTodo %s 更新成功", c.Params("id")),
				Data:       &data,
			})
		}
	}
	logrus.Error("user not found")
	return c.Status(fiber.StatusUnauthorized).JSON(dto.CommonResponse[any]{
		StatusCode: 401,
		InfoMsg:    "找不到該使用者: " + data.Username,
	})
}
