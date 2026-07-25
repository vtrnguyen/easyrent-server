package database

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"easyrent-server/internal/constants"
	"easyrent-server/internal/dto/requests"
)

// ApplyFilters applies the provided filters to the GORM database query based on the field mapping.
func ApplyFilters(
	db *gorm.DB,
	filters []requests.Filter,
	fieldMap map[string]string,
	filterLogic constants.FilterLogic,
) *gorm.DB {
	if len(filters) == 0 {
		return db
	}

	var conditions []string
	var values []interface{}

	for _, filter := range filters {
		column, exists := fieldMap[filter.Field]

		if !exists {
			continue
		}

		switch filter.Operator {
		case constants.FilterOperatorEquals:
			conditions = append(conditions, fmt.Sprintf("%s = ?", column))
			values = append(values, filter.Value)
		case constants.FilterOperatorContains:
			conditions = append(conditions, fmt.Sprintf("%s LIKE ?", column))
			values = append(values, fmt.Sprintf("%%%v%%", filter.Value))
		case constants.FilterOperatorIn:
			conditions = append(conditions, fmt.Sprintf("%s IN ?", column))
			values = append(values, filter.Value)
		}
	}

	if len(conditions) == 0 {
		return db
	}

	operator := fmt.Sprintf(" %s ", constants.FilterLogicAnd)

	if filterLogic == constants.FilterLogicOr {
		operator = fmt.Sprintf(" %s ", constants.FilterLogicOr)
	}

	return db.Where(strings.Join(conditions, operator), values...)
}

// ApplySorts applies the provided sorting criteria to the GORM database query based on the field mapping.
func ApplySorts(
	db *gorm.DB,
	sorts []requests.Sort,
	fieldMap map[string]string,
) *gorm.DB {
	for _, sort := range sorts {
		column, exists := fieldMap[sort.Field]
		if !exists {
			continue
		}

		direction := constants.SortOrderAsc
		if strings.ToLower(sort.Direction) == "desc" {
			direction = constants.SortOrderDesc
		}

		db = db.Order(
			fmt.Sprintf("%s %s", column, direction),
		)
	}

	return db
}
