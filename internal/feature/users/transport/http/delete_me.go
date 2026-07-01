package users_transport_http

import (
	"errors"
	"net/http"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteMe(w http.ResponseWriter, r *http.Request) {

	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		core_http_response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.svc.DeleteMe(r.Context(), userID); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			core_http_response.Error(w, http.StatusNotFound, "not found")
			return
		}
		core_http_response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	core_http_response.JSON(w, http.StatusNoContent, "user succesfully deleted")
}
