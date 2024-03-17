package user

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/user/database"
)

type UserHandler struct {
	storage *database.UserStore
}

func NewUserHandler(storage *database.UserStore) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (handler *UserHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	handler.storage.GetDiscount()
}
