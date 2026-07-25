package requests

import "easyrent-server/internal/constants"

type SearchRequest struct {
	Page        int                   `json:"page"`
	Limit       int                   `json:"limit"`
	FilterLogic constants.FilterLogic `json:"filter_logic"`
	Filters     []Filter              `json:"filters"`
	Sorts       []Sort                `json:"sorts"`
}

type Filter struct {
	Field    string                   `json:"field"`
	Operator constants.FilterOperator `json:"operator"`
	Value    interface{}              `json:"value"`
}

type Sort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}
