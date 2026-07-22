package validator

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
)

type TestStruct struct {
	Amount float64 `validate:"gt=0"`
	Gram   float64 `validate:"gt=0"`
}

func TestFormatValidationError(t *testing.T) {
	v := validator.New()

	t.Run("Test Required Field", func(t *testing.T) {
		req := TestStruct{Amount: 0}
		err := v.Struct(req)
		
		formattedErr := FormatValidationError(err)
		expected := "amount must be greater than zero"
		
		if formattedErr != expected {
			t.Errorf("Expected '%s', got '%s'", expected, formattedErr)
		}
	})

	t.Run("Test Negative Value", func(t *testing.T) {
		req := TestStruct{Amount: -500}
		err := v.Struct(req)
		
		formattedErr := FormatValidationError(err)
		expected := "amount must be greater than zero"
		
		if formattedErr != expected {
			t.Errorf("Expected '%s', got '%s'", expected, formattedErr)
		}
	})

	t.Run("Test Not Validation Error", func(t *testing.T) {
		standardErr := errors.New("internal server error")
		formattedErr := FormatValidationError(standardErr)
		
		if formattedErr != "internal server error" {
			t.Errorf("Expected 'internal server error', got '%s'", formattedErr)
		}
	})
}
