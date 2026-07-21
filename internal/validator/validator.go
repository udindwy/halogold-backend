package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			field := strings.ToLower(e.Field())
			tag := e.Tag()

			switch tag {
			case "required":
				return fmt.Sprintf("%s is required", field)
			case "gt":
				if e.Param() == "0" {
					return fmt.Sprintf("%s must be greater than zero", field)
				}
				return fmt.Sprintf("%s must be greater than %s", field, e.Param())
			default:
				return fmt.Sprintf("%s is invalid", field)
			}
		}
	}

	return err.Error()
}
