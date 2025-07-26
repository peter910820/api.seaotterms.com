package blog

import (
	"os"
	"time"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	// 因為驗證身份會去不同站台的資料庫抓該站台的身分驗證表，所以將該站台的資料庫寫死在middleware這邊，就不需要再路由額外傳遞
	authDbName = os.Getenv("DATABASE_NAME3")
)

type UserData struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Exp        int       `json:"exp"`
	Management bool      `json:"management"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdateName string    `json:"update_name"`
	Avatar     string    `json:"avatar"`
}

func CheckLogin(store *session.Store, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

func CheckOwner(store *session.Store, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
