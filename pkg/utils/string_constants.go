package utils

// Define a custom type for the context key
type contextKey string

const (
	EmailSubject            string     = "Sign Up OTP"
	EmailContent            string     = "Welcome to Gestapo !!!. Use the following OTP to complete your Sign Up procedures. OTP is valid for 5 minutes"
	Unauthorized            string     = "Unauthorized"
	AuthorizationKey        string     = "Authorization"
	AuthorizationTypeBearer string     = "bearer"
	AuthorizationPayloadKey contextKey = "authorization_payload"
)
