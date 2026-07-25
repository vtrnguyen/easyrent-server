package utils

import "strings"

// toSnakeCase converts a string to snake_case format.
func toSnakeCase(str string) string {
	var result strings.Builder

	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}

		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}
