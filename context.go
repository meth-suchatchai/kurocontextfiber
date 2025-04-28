package kuroctxfiber

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Context struct {
	*fiber.Ctx
}

func NewContext(ctx *fiber.Ctx) *Context {
	return &Context{ctx}
}

func (c *Context) BindAndValidate(params any) error {
	if err := c.BodyParser(params); err != nil {
		return err
	}

	if err := validate.Struct(params); err != nil {
		return err
	}

	return nil
}

func (c *Context) sendError(params any) error {
	return nil
}
