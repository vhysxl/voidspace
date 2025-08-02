package handler

import (
	"encoding/json"
	"net/http"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"
	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Credential string `json:"credential" validate:"required,min=3,max=50"`
	Password   string `json:"password" validate:"required,min=8"`
}

type LoginHandler struct {
	LoginUsecase usecase.LoginUsecase
	Validator    *validator.Validate
}

func NewLoginHandler(loginUsecase usecase.LoginUsecase, validator *validator.Validate) *LoginHandler {
	return &LoginHandler{
		LoginUsecase: loginUsecase,
		Validator:    validator,
	}
}

func (lh *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	err = lh.Validator.Struct(request)
	if err != nil {
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	_, err = lh.LoginUsecase.Login(r.Context(), request.Credential, request.Password)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			response.JSONErr(w, http.StatusUnauthorized, err.Error())
			return
		}
		response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	response.JSONSuccess[any](w, http.StatusOK, "Login successfully", nil)
}
