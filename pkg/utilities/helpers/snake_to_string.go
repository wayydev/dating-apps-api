package helpers

import "strings"

func SnackToString(value string) string {
	return strings.ReplaceAll(value, "_", " ")
}
