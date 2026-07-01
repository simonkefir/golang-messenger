package users_transport_http

import (
	"errors"
	"net/http"
	"strconv"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		core_http_response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	user, err := h.svc.GetUser(r.Context(), userID)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			core_http_response.Error(w, http.StatusNotFound, "not found")
			return
		}
		core_http_response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	core_http_response.JSON(w, http.StatusOK, NewUserResponseFromDomain(&user))
}
