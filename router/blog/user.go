package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
	middleware "api.seaotterms.com/middleware/blog"
)

func userRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string, store *session.Store) {
	userGroup := blogGroup.Group("/users")

	userGroup.Post("/", func(c *fiber.Ctx) error {
		return api.CreateUser(c, dbs[dbName])
	})
	userGroup.Patch("/:id", middleware.CheckLogin(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.UpdateUser(c, dbs[dbName])
	})
}
