package users_transport_http

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto CreateUserDTO

	if err := core_http_request.DecodeJSON(r, &dto); err != nil {
		log.Println("decode error:", err)
		core_http_response.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := h.validate.Struct(dto); err != nil {
		log.Println("validation error:", err)
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
		log.Println("service error:", err)
		if errors.Is(err, core_errors.ErrAlreadyExists) {
			core_http_response.Error(w, http.StatusConflict, "user already exists")
			return
		}
		core_http_response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	core_http_response.JSON(w, http.StatusCreated, NewUserResponseFromDomain(&user))
}
