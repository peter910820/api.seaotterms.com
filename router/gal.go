package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	galapi "api.seaotterms.com/api/gal"
)

func GalRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	galGroup := apiGroup.Group("/gal")

	authRouter(galGroup, dbs)
}

func authRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Post("/register", func(c *fiber.Ctx) error {
		return galapi.Register(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}
