package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("user_identity", isValidUserIdentity)
}

func isValidUserIdentity(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// Check if it's a valid email
	if validator.New().Var(username, "email") == nil {
		return true
	}

	// Check if it's alphanumeric
	if validator.New().Var(username, "alphanum,min=4,max=20") == nil {
		return true
	}

	return false
}

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = customErrorMessage(err)
	}
	return errors
}

func customErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("This field must be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("This field must be at most %s characters long", fe.Param())
	case "email":
		return "Invalid email format"
	case "alphanum":
		return "This field must be contain only alphanumeric characters"
	case "user_identity":
		return "This field must be valid email or username"
	default:
		return "Invalid field"
	}
}
