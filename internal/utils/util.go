package utils

import (
	"math"
	"strings"
)

// ToSnakeCase converts a string to snake_case format.
func ToSnakeCase(str string) string {
	var result strings.Builder

	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}

		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}

// CalculateTotalPages calculates the total number of pages based on the total number of items and the limit per page.
func CalculateTotalPages(total int64, limit int) int {
	if limit <= 0 {
		return 0
	}

	return int(
		math.Ceil(
			float64(total) / float64(limit),
		),
	)
}
