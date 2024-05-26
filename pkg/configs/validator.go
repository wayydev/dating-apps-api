package configs

import (
	"dating-apps/api/cmd/rules"
	"dating-apps/api/pkg/utilities/helpers"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"gorm.io/gorm"
)

var Translator ut.Translator

func InitValidate(db *gorm.DB) *validator.Validate {
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

	InitRule(validate, db)
	InitRuleMessage(validate, trans)

	return validate
}

func InitRule(validate *validator.Validate, db *gorm.DB) {
	validate.RegisterValidation("unique", rules.UniqueRule(db))
	validate.RegisterValidation("date", rules.DateRule)
}

func InitRuleMessage(validate *validator.Validate, trans ut.Translator) {
	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, helpers.SnackToString(fe.Field()), param)
		if err != nil {
			return fe.(error).Error()
		}

		return t
	}

	validate.RegisterTranslation("date", trans, rules.DateRuleMessage, transFn)
	validate.RegisterTranslation("unique", trans, rules.UniqueRuleMessage, transFn)
}
