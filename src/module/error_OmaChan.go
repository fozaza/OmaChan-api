package module

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorOmaChan struct {
	Err    error
	Status int
}

// set erros new
func New_ErrorOmChan() ErrorOmaChan {
	return ErrorOmaChan{nil, fiber.StatusOK}
}

// set errors new with status bad request
func (my_error ErrorOmaChan) Errors(msg string) ErrorOmaChan {
	my_error.Err = errors.New(msg)
	my_error.Status = fiber.StatusBadRequest
	return my_error
}

// / set status
func (my_error ErrorOmaChan) SetStatus(status int) ErrorOmaChan {
	my_error.Status = status
	return my_error
}

// set error bad BadServer
func (my_error ErrorOmaChan) BadServer() ErrorOmaChan {
	my_error.Status = fiber.StatusInternalServerError
	my_error.Err = errors.New("Internal server Error")
	return my_error
}

// set message errors
func (my_error ErrorOmaChan) SetErrorMsg(msg string) ErrorOmaChan {
	my_error.Err = errors.New(msg)
	return my_error
}

// for map ErrorOmaChan to fiber error
func (my_error ErrorOmaChan) MapFiber(c *fiber.Ctx) error {
	return c.Status(my_error.Status).SendString(my_error.Err.Error())
}
