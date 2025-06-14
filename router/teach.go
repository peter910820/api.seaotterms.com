package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"api.seaotterms.com/api"
)

func TeachRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	teachGroup := apiGroup.Group("/teach")

	SeriesRouter(teachGroup, dbs)
	ArticleRouter(teachGroup, dbs)
	CommentApiRouter(teachGroup, dbs)
}

func SeriesRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Get("/series", func(c *fiber.Ctx) error {
		return api.QuerySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Post("/series", func(c *fiber.Ctx) error {
		return api.CreateSeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Patch("/series/:id", func(c *fiber.Ctx) error {
		return api.ModifySeries(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}

func ArticleRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
	apiGroup.Get("/article", func(c *fiber.Ctx) error {
		return api.QueryArticle(c, dbs[os.Getenv("DATABASE_NAME")])
	})

	apiGroup.Post("/article", func(c *fiber.Ctx) error {
		return api.CreateArticle(c, dbs[os.Getenv("DATABASE_NAME")])
	})
}

func CommentApiRouter(apiGroup fiber.Router, dbs map[string]*gorm.DB) {
}
