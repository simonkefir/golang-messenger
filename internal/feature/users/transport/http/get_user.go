package users_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// GetUserByID      godoc
// @Summary     Получить пользователя
// @Description Получить пользователя в системе по ID
// @Tags        users
// @Produce     json
// @Param       id path int true "ID получаемого пользователя"
// @Success     200 {object} UserDTOResponse                  "Успешно полученный пользователь"
// @Failure     400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure     404 {object} core_http_response.ErrorResponse "Not found"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/{id}                                   [get]
func (h *UsersHTTPHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	idStr := r.PathValue("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid id")
		return
	}

	user, err := h.svc.GetUserByID(r.Context(), userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	responseHandler.JSONResponse(NewUserResponseFromDomain(&user), http.StatusOK)
}

// GetUserByUsername godoc
// @Summary          Получить пользователя
// @Description      Получить пользователя в системе по username
// @Tags             users
// @Produce          json
// @Param            username path string true                     "username получаемого пользователя"
// @Success          200 {object} UserDTOResponse                  "Успешно полученный пользователь"
// @Failure          400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure          404 {object} core_http_response.ErrorResponse "Not found"
// @Failure          500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router           /users/user/{username}                        [get]
func (h *UsersHTTPHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	username := r.PathValue("username")
	if username == "" {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid username")
		return
	}

	user, err := h.svc.GetUserByUsername(r.Context(), username)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	responseHandler.JSONResponse(NewUserResponseFromDomain(&user), http.StatusOK)
}
