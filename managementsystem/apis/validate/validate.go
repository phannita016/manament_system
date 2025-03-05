package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// create val instant validator.
var (
	val    = validator.New()
	AddNew = New{val: val}
)

// New struct implement validator
type New struct {
	val *validator.Validate
}

// Validation function validator body parser struct.
func Validation(i any) []*ErrorValidation {
	return validationFunc(val, i)
}

func (v New) Validate(i any) error {
	if err := validationFunc(v.val, i); len(err) > 0 {
		return ErrValidator(err)
	}
	return nil
}

func validationFunc(v *validator.Validate, i any) []*ErrorValidation {
	var validations []*ErrorValidation
	if err := v.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var element ErrorValidation
			element.Tag = e.Tag()
			element.Field = strings.ToLower(e.Field())
			element.Message = "Error:Field validation for '" + e.Field() + "' failed on the '" + e.Tag() + "' tag"
			validations = append(validations, &element)
		}
	}

	return validations
}
