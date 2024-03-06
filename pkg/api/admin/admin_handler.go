package admin

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/service/logger"
)

type AdminHandler struct {
	log logger.Logger
}

func NewAdminHandler(log logger.Logger) *AdminHandler {
	return &AdminHandler{
		log: log,
	}
}

func (admin *AdminHandler) CreateCategories(w http.ResponseWriter, r *http.Request) {

}
