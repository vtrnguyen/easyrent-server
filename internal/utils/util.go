package utils

import (
	"io"
	"math"
	"mime/multipart"
	"os"
	"strings"
	"time"
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

// ParseDate parses a date string in the format "YYYY-MM-DD" and returns a pointer to a time.Time object. If the input string is empty, it returns nil.
func ParseDate(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	date, err := time.Parse(
		"2006-01-02",
		value,
	)

	if err != nil {
		return nil, err
	}

	return &date, nil
}

// SaveFile saves the provided multipart file to the specified path on the filesystem.
func SaveFile(
	file *multipart.FileHeader,
	path string,
) error {
	src,err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	dst,err := os.Create(path)
	if err != nil {
		return err
	}

	defer dst.Close()
	_,err = io.Copy(dst,src)

	return err
}
