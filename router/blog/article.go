package blog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

func articleRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	articleGroup := apiGroup.Group("/article")

	articleGroup.Get("/", func(c *fiber.Ctx) error {
		return api.QueryArticle(c, dbs[dbName])
	})

	// No middleware has been implemented yet
	articleGroup.Post("/", func(c *fiber.Ctx) error {
		return api.CreateArticle(c, dbs[dbName])
	})
}
