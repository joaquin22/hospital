package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/patients/application"
)

type PatientHandler struct {
	createPatientUC *application.CreatePatientUseCase
	listPatientsUC  *application.ListPatientsUseCase
	getPatientUC    *application.GetPatientUseCase
	updatePatientUC *application.UpdatePatientUseCase
}

func NewPatientHandler(createPatientUC *application.CreatePatientUseCase, listPatientsUC *application.ListPatientsUseCase, getPatientUC *application.GetPatientUseCase, updatePatientUC *application.UpdatePatientUseCase) *PatientHandler {
	return &PatientHandler{
		createPatientUC: createPatientUC,
		listPatientsUC:  listPatientsUC,
		getPatientUC:    getPatientUC,
		updatePatientUC: updatePatientUC,
	}
}

type createPatientRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type UpdatePatientRequest struct {
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
	patients, err := h.listPatientsUC.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) GetPatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid patient ID"})
		return
	}

	patient, err := h.getPatientUC.GetPatient(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid patient ID"})
		return
	}

	var req UpdatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := h.updatePatientUC.Execute(application.UpdatePatientInput{
		ID:        uint(id),
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

	c.JSON(http.StatusOK, output)
}
