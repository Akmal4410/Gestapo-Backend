package merchent

import (
	"net/http"
)

type MerchentHandler struct {
}

func NewMerchentHandler() *MerchentHandler {
	return &MerchentHandler{}

}

func (merchet *MerchentHandler) ViewProfile(w http.ResponseWriter, r *http.Request) {
	// user_id := mux.Vars(r)["id"]

}
