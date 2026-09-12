package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
	"github.com/joaquin22/hospital-api/internal/users/application"
)

type UserHandler struct {
	registerUC  *application.RegisterUserUseCase
	loginUC     *application.LoginUserUseCase
	listUsersUC *application.ListUsersUseCase
}

func NewUserHandler(registerUC *application.RegisterUserUseCase, loginUC *application.LoginUserUseCase, listUsersUC *application.ListUsersUseCase) *UserHandler {
	return &UserHandler{
		registerUC:  registerUC,
		loginUC:     loginUC,
		listUsersUC: listUsersUC,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.registerUC.Execute(application.RegisterUserInput{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Dni:      req.Dni,
		Role:     req.Role,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error registering user", err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, output)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.loginUC.Execute(application.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error logging in", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.listUsersUC.Execute()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error listing users", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, users)
}
