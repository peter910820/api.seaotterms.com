package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	utils "api.seaotterms.com/utils/blog"
)

func AuthLogin(c *fiber.Ctx, store *session.Store) error {
	response := utils.ResponseFactory[any](c, fiber.StatusOK, "取得使用者資料成功", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}
