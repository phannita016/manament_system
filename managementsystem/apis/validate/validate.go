package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type ErrorValidation struct {
	Tag     string `json:"tag"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorValidator struct {
	Errors []*ErrorValidation `json:"errors"`
}

func (ev *ErrorValidator) Error() string {
	var errStrings []string
	for _, e := range ev.Errors {
		errStrings = append(errStrings, fmt.Sprintf("[%s]: %s", e.Field, e.Message))
	}
	return strings.Join(errStrings, ", ")
}

func Validate(i interface{}) error {
	err := validate.Struct(i)
	if err == nil {
		return nil
	}

	var validationErrors []*ErrorValidation
	for _, e := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, &ErrorValidation{
			Tag:     e.Tag(),
			Field:   strings.ToLower(e.Field()),
			Message: fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", e.Field(), e.Tag()),
		})
	}

	return &ErrorValidator{Errors: validationErrors}
}
