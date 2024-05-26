package rules

import (
	"fmt"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func DateRule(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := time.Parse("2006-01-02", date)
	fmt.Print(err)
	return err == nil
}

func DateRuleMessage(ut ut.Translator) error {
	return ut.Add("date", "{0} is not a valid date", false)
}
