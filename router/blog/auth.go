package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
	middleware "api.seaotterms.com/middleware/blog"
)

// this router is use to check identity for front-end routes
func AuthRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string, store *session.Store) {
	authGroup := blogGroup.Group("/auth")

	authGroup.Get("/", middleware.CheckLogin(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.AuthLogin(c, store)
	})
	// check if you are the website owner
	authGroup.Get("/root", middleware.CheckOwner(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.AuthLogin(c, store)
	})
	// authGroup.Get("/specific", middleware.CheckLogin(store, dbs[dbName]), func(c *fiber.Ctx) error {
	// 	return api.AuthLogin(c, store)
	// })
}
