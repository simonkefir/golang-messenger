package users_transport_http

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// Login        godoc
// @Summary     Залогиниться
// @Description Получить JWT-токен для авторизации себя
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body LoginUserDTO true "LoginUser тело запроса".
// @Success     200 {object} LoginResponse "Успешная аутентификация. Возращает JWT-токен"
// @Failure     400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/login [post]
func (h *UsersHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var dto LoginUserDTO

	if err := core_http_request.DecodeJSON(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "invalid json")
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

	token, err := h.svc.Login(r.Context(), dto.Email, dto.Password)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "failed to login")
		return
	}

	responseHandler.JSONResponse(LoginResponse{Token: token}, http.StatusOK)
}
