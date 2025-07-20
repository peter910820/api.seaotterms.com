package blog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

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
