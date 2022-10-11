package validations

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateTitleIsCool(field validator.FieldLevel) bool {
	value := strings.ToLower(field.Field().String())
	return strings.Contains(value, "cool")
}
