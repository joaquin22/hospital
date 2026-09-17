package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
	"github.com/joaquin22/hospital-api/internal/users/application"
)

type UserHandler struct {
	registerUC   *application.RegisterUserUseCase
	loginUC      *application.LoginUserUseCase
	listUsersUC  *application.ListUsersUseCase
	getUserUC    *application.GetUserUseCase
	updateUserUC *application.UpdateUserUseCase
	patchUserUC  *application.PatchUserUseCase
}

func NewUserHandler(registerUC *application.RegisterUserUseCase, loginUC *application.LoginUserUseCase, listUsersUC *application.ListUsersUseCase, getUserUC *application.GetUserUseCase, updateUserUC *application.UpdateUserUseCase, patchUserUC *application.PatchUserUseCase) *UserHandler {
	return &UserHandler{
		registerUC:   registerUC,
		loginUC:      loginUC,
		listUsersUC:  listUsersUC,
		getUserUC:    getUserUC,
		updateUserUC: updateUserUC,
		patchUserUC:  patchUserUC,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.registerUC.Execute(application.RegisterUserInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Dni:       req.Dni,
		Role:      req.Role,
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
		response.ErrorResponse(c, http.StatusBadRequest, "error logging in", err)
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

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid user ID", err)
		return
	}

	output, err := h.getUserUC.Execute(uint(id))

	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "error getting user", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid user ID", err)
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.updateUserUC.Execute(application.UpdateUserInput{
		ID:        uint(id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Dni:       req.Dni,
		Role:      req.Role,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error updating user", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *UserHandler) PatchUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid user ID", err)
		return
	}

	var req patchUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.patchUserUC.Execute(application.PatchUserInput{
		ID:        uint(id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Dni:       req.Dni,
		Role:      req.Role,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error patching user", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}
