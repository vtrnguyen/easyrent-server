package constants

type UserStatus string

type Gender string

type AccountRole string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"

	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"

	AccountRoleAdmin    AccountRole = "admin"
	AccountRoleTenant   AccountRole = "tenant"
	AccountRoleLandlord AccountRole = "landlord"
)
