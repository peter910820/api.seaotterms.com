package blog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

func articleRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	articleGroup := blogGroup.Group("/articles")

	articleGroup.Get("/", func(c *fiber.Ctx) error {
		return api.QueryArticle(c, dbs[dbName])
	})

	articleGroup.Get("/:id", func(c *fiber.Ctx) error {
		return api.QueryArticle(c, dbs[dbName])
	})

	// No middleware has been implemented yet
	articleGroup.Post("/", func(c *fiber.Ctx) error {
		return api.CreateArticle(c, dbs[dbName])
	})

	// articleGroup.Post("/:id", func(c *fiber.Ctx) error {
	// 	return api.ModifyArticle(c, dbs[dbName])
	// })

	articleGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return api.DeleteArticle(c, dbs[dbName])
	})
}
