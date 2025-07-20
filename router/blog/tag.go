package blog

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

func tagRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	apiGroup.Get("/tag", func(c *fiber.Ctx) error {
		return api.QueryTag(c, dbs[dbName])
	})

	apiGroup.Post("/tag", func(c *fiber.Ctx) error {
		return api.CreateTag(c, dbs[dbName])
	})
}
