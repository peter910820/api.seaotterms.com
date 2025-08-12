package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
	middleware "api.seaotterms.com/middleware/blog"
)

func GalgameBrandRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string, store *session.Store) {
	galgameBrandGroup := blogGroup.Group("/galgame-brands")

	galgameBrandGroup.Get("/", func(c *fiber.Ctx) error {
		return api.QueryAllGalgameBrand(c, dbs[dbName])
	})
	galgameBrandGroup.Get("/:brand", func(c *fiber.Ctx) error {
		return api.QueryGalgameBrand(c, dbs[dbName])
	})
	galgameBrandGroup.Post("/", middleware.CheckManagement(store), func(c *fiber.Ctx) error {
		return api.CreateGalgameBrand(c, dbs[dbName])
	})
	galgameBrandGroup.Patch("/:brand", middleware.CheckManagement(store), func(c *fiber.Ctx) error {
		return api.UpdateGalgameBrand(c, dbs[dbName])
	})
}
