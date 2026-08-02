package constants

type UserStatus string
type Gender string
type AccountRole string
type PropertyType string
type SortOrder string
type FilterOperator string
type FilterLogic string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"

	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"

	AccountRoleAdmin    AccountRole = "admin"
	AccountRoleTenant   AccountRole = "tenant"
	AccountRoleLandlord AccountRole = "landlord"

	PropertyTypeHouse      PropertyType = "house"
	PropertyTypeRentalRoom PropertyType = "rental_room"
	PropertyTypeApartment  PropertyType = "apartment"
	PropertyTypeFlat       PropertyType = "flat"

	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"

	FilterOperatorEquals   FilterOperator = "equals"
	FilterOperatorContains FilterOperator = "contains"
	FilterOperatorIn       FilterOperator = "in"

	FilterLogicAnd FilterLogic = "and"
	FilterLogicOr  FilterLogic = "or"
)
