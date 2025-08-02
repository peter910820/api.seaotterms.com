package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
)

func loginRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string, store *session.Store) {
	loginGroup := blogGroup.Group("/login")

	loginGroup.Post("/", func(c *fiber.Ctx) error {
		return api.Login(c, store, dbs[dbName])
	})
}
