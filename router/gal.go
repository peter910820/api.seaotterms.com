package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	galapi "api.seaotterms.com/api/gal"
)

func GalRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	galGroup := apiGroup.Group("/gal")
	dbName := os.Getenv("DATABASE_NAME2")

	authRouter(galGroup, dbs, dbName)
	loginRouter(galGroup, dbs, dbName)
	testRouter(galGroup)
}

func authRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	apiGroup.Post("/register", func(c *fiber.Ctx) error {
		return galapi.Register(c, dbs[dbName])
	})

	apiGroup.Get("/register/:register-key", func(c *fiber.Ctx) error {
		return galapi.RegisterKeyCheck(c, dbs[dbName])
	})
}

func loginRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB, dbName string) {
	apiGroup.Post("/login", func(c *fiber.Ctx) error {
		return galapi.Register(c, dbs[dbName])
	})
}

func testRouter(apiGroup fiber.Router) {
	apiGroup.Get("/test-mail-send", func(c *fiber.Ctx) error {
		return galapi.TestMailSend(c)
	})
}
