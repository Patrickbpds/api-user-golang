package validation

import (
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transL   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transL, _ = uni.GetTranslator("en")
		en_translator.RegisterDefaultTranslations(val, transL)
	}
}

func ValidateUserError(
	validation_errors error,
) *rest_errors.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_errors, &jsonErr) {
		return rest_errors.NewBadRequestError(
			"Invalid fild format: " + jsonErr.Field)
	} else if errors.As(validation_errors, &jsonValidationError) {
		errorsCauses := []rest_errors.Causes{}

		for _, e := range validation_errors.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Field:   e.Field(),
				Message: e.Translate(transL),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_errors.NewBadRequestError(
			"Error trying to convert fields: " + validation_errors.Error())
	}
}
