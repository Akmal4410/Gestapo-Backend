package utils

// constants for all supported SignupAction
const (
	SIGN_UP         string = "sign-up"
	FORGOT_PASSWORD string = "forgot-password"
)

// IsSupportedSignupAction returns true if the usertype is supported
func IsSupportedSignupAction(action string) bool {
	switch action {
	case SIGN_UP, FORGOT_PASSWORD:
		return true
	}
	return false
}
