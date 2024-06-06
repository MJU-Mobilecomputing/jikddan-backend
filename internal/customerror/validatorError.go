package customerror

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidatorError struct {
	Key string
	Msg string
}

type ValidatorErrors struct {
	Errors []ValidatorError
}

var VALIDATOR_ERROR = NewError(http.StatusBadRequest, "param validate fail.", "001")

func ValidateError(err error) error {
	validatorError := VALIDATOR_ERROR
	var errors ValidatorErrors
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErr {
			v := ValidatorError{
				Key: e.Field(),
				Msg: e.Error(),
			}
			errors.Errors = append(errors.Errors, v)
		}
	}
	validatorError.Data = errors
	return validatorError
}
