package configs

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var Translator ut.Translator

func InitValidate() *validator.Validate {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("json")
		if name != "" {
			return name
		}
		return field.Name
	})

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	enTranslations.RegisterDefaultTranslations(validate, trans)
	Translator = trans

	return validate
}
