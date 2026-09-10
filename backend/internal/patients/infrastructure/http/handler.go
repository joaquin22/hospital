package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/patients/application"
)

type PatientHandler struct {
	createPatientUC *application.CreatePatientUseCase
	listPatientsUC  *application.ListPatientsUseCase
}

func NewPatientHandler(createPatientUC *application.CreatePatientUseCase, listPatientsUC *application.ListPatientsUseCase) *PatientHandler {
	return &PatientHandler{
		createPatientUC: createPatientUC,
		listPatientsUC:  listPatientsUC,
	}
}

type createPatientRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var req createPatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := h.createPatientUC.Execute(application.CreatePatientInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Dni:       req.Dni,
		Phone:     req.Phone,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *PatientHandler) ListPatients(c *gin.Context) {
	patients, err := h.listPatientsUC.ListPatients()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}
