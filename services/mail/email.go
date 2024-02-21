package mail

import "github.com/akmal4410/gestapo/services/cache"

type EmailService interface {
	SendOTP(to, subject, content string, redis cache.Cache) error
	VerfiyOTP(to, code string, redis cache.Cache) (bool, error)
}
