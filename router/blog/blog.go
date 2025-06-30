package blog

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

func BlogRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	blogGroup := apiGroup.Group("/blog")
	dbName := os.Getenv("DATABASE_NAME3")

	articleRouter(blogGroup, dbs, dbName)
	tagRouter(blogGroup, dbs, dbName)
}

func articleRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	apiGroup.Get("/article", func(c *fiber.Ctx) error {
		return api.QueryArticle(c, dbs[dbName])
	})

	// No middleware has been implemented yet
	apiGroup.Post("/article", func(c *fiber.Ctx) error {
		return api.CreateArticle(c, dbs[dbName])
	})

	apiGroup.Get("/article/:articleID", func(c *fiber.Ctx) error {
		return api.QuerySingleArticle(c, dbs[dbName])
	})
}

func tagRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	apiGroup.Get("/tag", func(c *fiber.Ctx) error {
		return api.QueryTag(c, dbs[dbName])
	})

	apiGroup.Post("/tag", func(c *fiber.Ctx) error {
		return api.CreateTag(c, dbs[dbName])
	})
}
