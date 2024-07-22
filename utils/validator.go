package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate
var Trans ut.Translator

func ValidateStruct(s interface{}) error {
	if err := Validate.Struct(s); err != nil {
		return err
	}
	return nil
}
func TranslateValidationError(err error) string {
	errs := err.(validator.ValidationErrors)
	var errMsg string
	for _, e := range errs {
		errMsg += e.Translate(Trans) + ","
	}
	return errMsg
}
