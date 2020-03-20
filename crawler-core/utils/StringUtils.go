package utils

import "strings"

func ReplaceAllEmpty(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), " ", "")
}
