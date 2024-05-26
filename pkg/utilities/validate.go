package utilities

import (
	"dating-apps/api/pkg/configs"

	"github.com/go-playground/validator/v10"
)

func TranslatorValidate(err error) map[string]string {
	errs := make(map[string]string)
	if err != nil {
		validatorErrors := err.(validator.ValidationErrors)
		for _, e := range validatorErrors {
			translated := e.Translate(configs.Translator)
			errs[e.Field()] = translated
		}
		return errs
	}

	return nil
}
