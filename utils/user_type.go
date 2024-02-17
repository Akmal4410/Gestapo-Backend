package utils

// constants for all supported user_type
const (
	USER     = "USER"
	MERCHANT = "MERCHANT"
	ADMIN    = "ADMIN"
)

// IsSupportedUsers returns true if the usertype is supported
func IsSupportedUsers(usertType string) bool {
	switch usertType {
	case USER, MERCHANT, ADMIN:
		return true
	}
	return false
}
