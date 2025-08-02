package blog

import (
	"errors"
	"os"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	dto "api.seaotterms.com/dto/blog"
	utils "api.seaotterms.com/utils/blog"
)

var (
	// 因為驗證身份會去不同站台的資料庫抓該站台的身分驗證表，所以將該站台的資料庫寫死在middleware這邊，就不需要再路由額外傳遞
	authDbName string
	// 建立共用使用者資料表，用來處理同使用者登入不同瀏覽器的狀況
	UserInfo = map[uint]*dto.UserInfo{}
)

func init() {
	// init env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file load error: %v", err)
	}
	authDbName = os.Getenv("DATABASE_NAME3")
}

// 用Token檢查使用者資料(預設全域註冊、回傳)
func GetUserInfo(store *session.Store, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userInfo, _ := checkLogin(c, store)
		c.Locals("user_info", userInfo)
		return c.Next()
	}
}

// 檢查有無登入
func CheckLogin(store *session.Store, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userInfo, err := checkLogin(c, store)
		if err != nil {
			logrus.Warn(err)
			response := utils.ResponseFactory[any](c, fiber.StatusUnauthorized, err.Error(), nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}
		c.Locals("user_info", userInfo)
		return c.Next()
	}
}

// 檢查是不是網站管理者
func CheckOwner(store *session.Store, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

// utils
func checkLogin(c *fiber.Ctx, store *session.Store) (*dto.UserInfo, error) {
	sess, err := store.Get(c)
	if err != nil {
		logrus.Fatal(err)
	}
	userID := sess.Get("id")
	if userID == nil {
		return nil, errors.New("visitors is not logged in")
	}

	userInfo, ok := UserInfo[userID.(uint)]
	if !ok {
		return nil, errors.New("visitors is not logged in")
	}

	return userInfo, nil
}
