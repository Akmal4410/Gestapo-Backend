package user

import "net/http"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (handler *UserHandler) GetHome(w http.ResponseWriter, r *http.Request) {

}
