package validate

import (
	"log/slog"
	"reflect"

	"github.com/labstack/echo/v4"
)

type TypeParserFunc func(msg string, args ...any)

func BodyParser[T any](c echo.Context, fps ...TypeParserFunc) (*T, error) {
	var t T

	var logs = func(msg string, body any, err any) {
		for _, fp := range fps {
			fp("BodyParser-"+msg,
				slog.String("struct", reflect.TypeOf(t).Name()),
				slog.Any("body", body),
				slog.Any("err", err),
			)
		}
	}

	if err := c.Bind(&t); err != nil {
		defer logs("Bind", nil, err)
		return nil, err
	}

	if err := c.Validate(&t); err != nil {
		defer logs("Context-Validation", t, err)
		return nil, err
	}

	return &t, nil
}
