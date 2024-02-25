package utils

// constants for all supported SSO Action
const (
	SSO_ANDROID string = "sso-android"
	SSO_IOS     string = "sso-ios"
)

// IsSupportedSSOAction returns true if the usertype is supported
func IsSupportedSSOAction(action string) bool {
	switch action {
	case SSO_ANDROID, SSO_IOS:
		return true
	}
	return false
}
