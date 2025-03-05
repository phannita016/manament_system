package types

type APIsHTTPError struct {
	Code        string `json:"code"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Validations any    `json:"validations,omitempty"`
}
