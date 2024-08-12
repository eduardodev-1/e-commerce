package httpError

import "github.com/gofiber/fiber/v2"

type HttpCustomError struct {
	Ctx *fiber.Ctx
}

func (f *HttpCustomError) NewHttpError(params *ErrorParams) error {
	return f.Ctx.Status(params.Status).JSON(fiber.Map{
		"error": params.Message,
	})
}

type ErrorParams struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (p *ErrorParams) SetDefaultParams(err error) {
	p.Status = fiber.StatusInternalServerError
	p.Message = err.Error()
}

func (p *ErrorParams) SetCustomError(status int, message string) {
	p.Status = status
	p.Message = message
}
