package users_transport_http

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// PatchMe      godoc
// @Summary     Изменить себя
// @Description Изменить своего пользователя в системе
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body PatchUserDTO true                "PatchMe тело запроса"
// @Security    BearerAuth
// @Success     200 {object} UserDTOResponse                  "Успешно изменённый пользователь"
// @Failure     400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/me                                     [patch]
func (h *UsersHTTPHandler) PatchMe(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var dto PatchUserDTO

	if err := core_http_request.DecodeJSON(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "invalid json")
		return
	}

	if dto.Username == nil && dto.DisplayName == nil && dto.Email == nil && dto.Password == nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "no fields to change")
		return
	}

	if err := h.validate.Struct(dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make(map[string]string)
			for _, fe := range ve {
				errs[fe.Field()] = fe.Tag()
			}
			core_http_response.ValidationError(w, errs)
			return
		}
	}

	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	user, err := h.svc.PatchMe(r.Context(), userID, dto.Username, dto.DisplayName, dto.Password, dto.Email)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update user")
		return
	}

	responseHandler.JSONResponse(NewUserResponseFromDomain(&user), http.StatusOK)
}
