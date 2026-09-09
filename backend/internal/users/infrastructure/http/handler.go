package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/users/application"
)

type UserHandler struct {
	registerUC *application.RegisterUserUseCase
	loginUC    *application.LoginUserUseCase
}

func NewUserHandler(registerUC *application.RegisterUserUseCase, loginUC *application.LoginUserUseCase) *UserHandler {
	return &UserHandler{
		registerUC: registerUC,
		loginUC:    loginUC,
	}
}

type registerRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Dni      string `json:"dni" binding:"required"`
	Role     string `json:"role" binding:"required"` // ADMIN o RECEPTIONIST
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
