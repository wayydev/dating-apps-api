package helpers

import (
	"strings"
	"unicode"
)

func CamelToSnake(s string) string {
	var result strings.Builder
	result.Grow(len(s) + 5)

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}

	return result.String()
}
