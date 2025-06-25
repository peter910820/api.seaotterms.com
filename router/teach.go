package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	teachapi "api.seaotterms.com/api/teach"
)

func TeachRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	teachGroup := apiGroup.Group("/teach")

	seriesRouter(teachGroup, dbs)
	articleRouter(teachGroup, dbs)
	commentApiRouter(teachGroup, dbs)
}

func seriesRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Get("/series", func(c *fiber.Ctx) error {
		return teachapi.QuerySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Post("/series", func(c *fiber.Ctx) error {
		return teachapi.CreateSeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Patch("/series/:id", func(c *fiber.Ctx) error {
		return teachapi.ModifySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}

func articleRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Get("/article", func(c *fiber.Ctx) error {
		return teachapi.QueryArticle(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Post("/article", func(c *fiber.Ctx) error {
		return teachapi.CreateArticle(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}

func commentApiRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
}
