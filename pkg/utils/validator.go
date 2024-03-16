package utils

// constants for all supported SignupAction
const (
	SIGN_UP         string = "sign-up"
	FORGOT_PASSWORD string = "forgot-password"
)

// constants for all supported SSO Action
const (
	SSO_ANDROID string = "sso-android"
	SSO_IOS     string = "sso-ios"
)

// constants for all supported USER_TYPE
const (
	USER     = "USER"
	MERCHANT = "MERCHANT"
	ADMIN    = "ADMIN"
)

// constants for all supported Gender
const (
	MALE   = "Male"
	FEMALE = "Female"
)

// IsSupportedSignupAction returns true if the SignupAction is supported
func IsSupportedSignupAction(action string) bool {
	switch action {
	case SIGN_UP, FORGOT_PASSWORD:
		return true
	}
	return false
}

// IsSupportedSSOAction returns true if the SSOAction is supported
func IsSupportedSSOAction(action string) bool {
	switch action {
	case SSO_ANDROID, SSO_IOS:
		return true
	}
	return false
}

// IsSupportedUsers returns true if the USER_TYPE is supported
func IsSupportedUsers(usertType string) bool {
	switch usertType {
	case USER, MERCHANT, ADMIN:
		return true
	}
	return false
}

// IsSupportedUsers returns true if the Gender is supported
func IsSupportedGender(gender string) bool {
	switch gender {
	case MALE, FEMALE:
		return true
	}
	return false
}

// IsSupportedPercentage returns true if the percentage is supported
func IsSupportedPercentage(percentage float64) bool {
	if percentage > 1 && percentage < 100 {
		return true
	}
	return false
}
