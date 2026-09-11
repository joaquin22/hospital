package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(400, gin.H{"error": err.Error()})
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
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := h.loginUC.Execute(application.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.listUsersUC.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
