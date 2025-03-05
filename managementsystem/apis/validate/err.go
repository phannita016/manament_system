package validate

import (
	"encoding/json"
	"net/http"
)

type (
	ErrorValidator struct {
		Code           int
		Opt            string
		ErrorResponses []*ErrorValidation
	}

	ErrorValidation struct {
		Tag     string `json:"tag"`
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func (e *ErrorValidator) Error() string {
	if e == nil {
		return "Unknown nil ErrorValidator"
	}

	bt, err := json.Marshal(e.ErrorResponses)
	if err != nil {
		return err.Error()
	}

	return string(bt)
}

func (e *ErrorValidator) ErrMessage() string {
	return "validation tag error please check payload"
}

func ErrValidator(errs []*ErrorValidation) error {
	return &ErrorValidator{Code: http.StatusBadRequest, Opt: "VALIDATION", ErrorResponses: errs}
}
