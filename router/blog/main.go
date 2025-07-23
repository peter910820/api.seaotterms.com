package blog

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

// 除了身份驗證表的資料庫，其餘資料庫名稱都定義在各站台router包的main.go中
func BlogRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, store *session.Store) {
	blogGroup := apiGroup.Group("/blog")
	dbName := os.Getenv("DATABASE_NAME3")

	articleRouter(blogGroup, dbs, dbName)
	tagRouter(blogGroup, dbs, dbName)
	TodoRouter(blogGroup, dbs, dbName, store)
	SystemTodoRouter(blogGroup, dbs, dbName, store)
}
