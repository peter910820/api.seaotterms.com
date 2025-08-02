package blog

import (
	"github.com/gofiber/fiber/v2"

	dto "api.seaotterms.com/dto/blog"
)

func ResponseFactory[T any](c *fiber.Ctx, httpStatus int, msg string, data *T) dto.CommonResponse[T] {
	if httpStatus == 200 {
		return dto.CommonResponse[T]{
			StatusCode: httpStatus,
			InfoMsg:    msg,
			UserInfo:   c.Locals("user_info").(dto.UserInfo),
			Data:       data,
		}
	} else {
		return dto.CommonResponse[T]{
			StatusCode: httpStatus,
			ErrMsg:     msg,
			UserInfo:   c.Locals("user_info").(dto.UserInfo),
			Data:       data,
		}
	}
}
