package blog

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BlogRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	blogGroup := apiGroup.Group("/blog")
	dbName := os.Getenv("DATABASE_NAME3")

	articleRouter(blogGroup, dbs, dbName)
	tagRouter(blogGroup, dbs, dbName)
}
