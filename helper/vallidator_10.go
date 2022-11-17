package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("%s is %s", e.Field(), e.ActualTag()))
	}
	return errors
}

// func TranslateError(err error) []string {
// 	if err == nil {
// 		return nil
// 	}

// 	var errors []string

// 	validate := validator.New()
// 	english := en.New()
// 	uni := ut.New(english, english)
// 	trans, _ := uni.GetTranslator("en")

// 	en_translations.RegisterDefaultTranslations(validate, trans)

// 	validatorErrs := err.(validator.ValidationErrors)
// 	for _, e := range validatorErrs {
// 		translatedErr := string(e.Translate(trans))
// 		errors = append(errors, translatedErr)

// 	}

// 	return errors
// }

// func FormatValidationError(err error) string {
// 	var errors string
// 	for _, e := range err.(validator.ValidationErrors) {
// 		errors = fmt.Sprintf("Field is %s", e.ActualTag())
// 	}
// 	return errors
// }
