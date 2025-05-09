package kuroctxfiber

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var JWTConfig = jwtware.Config{}

func (f *defaultFiber) EnableJWTMiddleware(config ...jwtware.Config) {
	if len(config) != 0 {
		JWTConfig = config[0]
	}

	JWTConfig.ErrorHandler = func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	f.server.Use(jwtware.New(JWTConfig))
}
