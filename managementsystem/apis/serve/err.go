package serve

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/apis/validate"
	"github.com/phannita016/management/types"
)

// APIsHTTPError struct format response error.
type APIsHTTPError = types.APIsHTTPError

var (
	InternalError = APIsHTTPError{
		Code:    "InternalServerError",
		Status:  "ERROR",
		Message: http.StatusText(http.StatusInternalServerError),
	}
)

// HandleError function handling error.
// handling error from handler return.
func HandleError(err error, c echo.Context) {
	const msg = "HandleError"
	// committed response.
	if c.Response().Committed {
		return
	}

	var e *echo.HTTPError
	if ok := errors.As(err, &e); ok {
		var (
			txt    = http.StatusText(e.Code)
			splits = strings.Split(txt, " ")

			// variable val body response.
			val = APIsHTTPError{
				Code:    splits[len(splits)-1],
				Status:  txt,
				Message: fmt.Sprint(e.Message),
			}
		)
		if e.Internal != nil {
			var ev *validate.ErrorValidator
			if ok := errors.As(e.Internal, &ev); ok {
				e.Code = ev.Code
				val.Code = ev.Opt
				val.Status = http.StatusText(ev.Code)
				val.Message = ev.ErrMessage()
				val.Validations = ev.ErrorResponses
			}
		}

		// response http-error.
		if err = c.JSON(e.Code, val); err != nil {
			slog.ErrorContext(c.Request().Context(), msg, slog.Any("err", err))
		}
		// return end function HandleError.
		return
	}

	if err = c.JSON(http.StatusInternalServerError, InternalError); err != nil {
		slog.ErrorContext(c.Request().Context(), msg, slog.Any("err", err))
	}
}
