package validate

import "github.com/labstack/echo/v4"

type EchoValidator struct{}

func NewEchoValidator() *EchoValidator {
	return &EchoValidator{}
}

func (*EchoValidator) Validate(i interface{}) error {
	return Validate(i)
}

func RegisterValidator(e *echo.Echo) {
	e.Validator = NewEchoValidator()
}
