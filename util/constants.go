package util

import (
	"net/http"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error
