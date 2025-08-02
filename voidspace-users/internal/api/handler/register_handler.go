package handler

import (
	"encoding/json"
	"net/http"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"
	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterHandler struct {
	RegisterUsecase usecase.RegisterUsecase
	Validator       *validator.Validate
}

func NewRegisterHandler(RegisterUsecase usecase.RegisterUsecase, validator *validator.Validate) *RegisterHandler {
	return &RegisterHandler{
		RegisterUsecase: RegisterUsecase,
		Validator:       validator,
	}
}

func (rh *RegisterHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	err = rh.Validator.Struct(request)
	if err != nil {
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	hashed := string(hashedPassword)

	_, err = rh.RegisterUsecase.Register(r.Context(), request.Username, request.Email, hashed)
	if err != nil {
		switch err {
		case domain.ErrUserExists:
			response.JSONErr(w, http.StatusConflict, err.Error())
			return
		case domain.ErrEmailExists:
			response.JSONErr(w, http.StatusConflict, err.Error())
			return
		default:
			response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
			return
		}
	}

	response.JSONSuccess[any](w, http.StatusCreated, "User registered successfully", nil)
}
