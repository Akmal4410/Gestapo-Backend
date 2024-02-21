package utils

// constants for all supported USER_TYPE
const (
	USER     = "USER"
	MERCHANT = "MERCHANT"
	ADMIN    = "ADMIN"
)

// IsSupportedUsers returns true if the USER_TYPE is supported
func IsSupportedUsers(usertType string) bool {
	switch usertType {
	case USER, MERCHANT, ADMIN:
		return true
	}
	return false
}
