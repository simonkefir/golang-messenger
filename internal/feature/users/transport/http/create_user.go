package users_transport_http

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// CreateUser   godoc
// @Summary     Создать пользователя
// @Description Создать нового пользователя в системе
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body CreateUserDTO true "CreateUser тело запроса"
// @Success     201 {object} UserDTOResponse "Успешно созданный пользователь"
// @Failure     400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/register [post]
func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var dto CreateUserDTO

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

	user, err := h.svc.CreateUser(r.Context(), domain.User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: dto.Password,
	})
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responseHandler.JSONResponse(NewUserResponseFromDomain(&user), http.StatusCreated)
}
