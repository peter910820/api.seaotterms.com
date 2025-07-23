package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	api "api.seaotterms.com/api/blog"
	middleware "api.seaotterms.com/middleware/blog"
)

func SystemTodoRouter(blogGroup fiber.Router, dbs map[string]*gorm.DB, dbName string, store *session.Store) {
	systemTodoGroup := blogGroup.Group("/system-todos")

	systemTodoGroup.Get("/", func(c *fiber.Ctx) error {
		return api.QuerySystemTodo(c, dbs[dbName])
	})

	systemTodoGroup.Post("/", middleware.CheckLogin(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.CreateSystemTodo(c, dbs[dbName])
	})

	systemTodoGroup.Patch("/:id", middleware.CheckOwner(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.UpdateSystemTodo(c, dbs[dbName])
	})

	// quick update
	systemTodoGroup.Patch("/quick/:id", middleware.CheckOwner(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.QuickUpdateSystemTodo(c, dbs[dbName])
	})

	systemTodoGroup.Delete("/:id", middleware.CheckOwner(store, dbs[dbName]), func(c *fiber.Ctx) error {
		return api.DeleteSystemTodo(c, dbs[dbName])
	})
}
