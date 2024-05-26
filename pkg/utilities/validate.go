package utilities

import (
	"dating-apps/api/pkg/configs"
	"fmt"

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
		fmt.Println(errs)
		return errs
	}

	return nil
}
