package blog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

func tagRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	tagGroup := blogGroup.Group("/tag")

	tagGroup.Get("/", func(c *fiber.Ctx) error {
		return api.QueryTag(c, dbs[dbName])
	})

	tagGroup.Get("/:name", func(c *fiber.Ctx) error {
		return api.QueryArticleForTag(c, dbs[dbName])
	})

	tagGroup.Post("/", func(c *fiber.Ctx) error {
		return api.CreateTag(c, dbs[dbName])
	})
}
