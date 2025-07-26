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
	dbName2 := os.Getenv("DATABASE_NAME2") // galgame DB

	articleRouter(blogGroup, dbs, dbName)
	tagRouter(blogGroup, dbs, dbName)
	TodoRouter(blogGroup, dbs, dbName, store)
	SystemTodoRouter(blogGroup, dbs, dbName, store)
	AuthRouter(blogGroup, dbs, dbName, store)
	UserRouter(blogGroup, dbs, dbName, store)
	TodoTopicRouter(blogGroup, dbs, dbName, store)
	GalgameRouter(blogGroup, dbs, dbName2, store)
	GalgameBrandRouter(blogGroup, dbs, dbName2, store)
	brandRouter(blogGroup, dbs, dbName2, store)
}
