package rules

import (
	"dating-apps/api/pkg/utilities/helpers"
	"fmt"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UniqueRule(db *gorm.DB) validator.Func {
	return func(fl validator.FieldLevel) bool {
		field := fl.Field()
		params := strings.Split(fl.Param(), ",")

		if len(params) == 0 {
			panic("Select the table name on the unique validation")
		}

		table := params[0]

		var wherePlainSql string

		switch len(params) {
		case 2:
			wherePlainSql = strings.ReplaceAll(params[1], "{value}", field.String())
		default:
			wherePlainSql = fmt.Sprintf("%s = '%s'", helpers.CamelToSnake(fl.StructFieldName()), field.String())
		}

		var count int64
		if err := db.Table(table).Where(wherePlainSql).Count(&count).Error; err != nil {
			panic(err)
		}

		return count == 0
	}
}

func UniqueRuleMessage(ut ut.Translator) error {
	return ut.Add("unique", "{0} is already", false)
}
