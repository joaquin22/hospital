package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/patients/application"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
)

type PatientHandler struct {
	createPatientUC *application.CreatePatientUseCase
	listPatientsUC  *application.ListPatientsUseCase
	getPatientUC    *application.GetPatientUseCase
	updatePatientUC *application.UpdatePatientUseCase
	patchPatientUC  *application.PatchPatientUseCase
}

func NewPatientHandler(createPatientUC *application.CreatePatientUseCase, listPatientsUC *application.ListPatientsUseCase, getPatientUC *application.GetPatientUseCase, updatePatientUC *application.UpdatePatientUseCase, patchPatientUC *application.PatchPatientUseCase) *PatientHandler {
	return &PatientHandler{
		createPatientUC: createPatientUC,
		listPatientsUC:  listPatientsUC,
		getPatientUC:    getPatientUC,
		updatePatientUC: updatePatientUC,
		patchPatientUC:  patchPatientUC,
	}
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var req createPatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
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
		response.ErrorResponse(c, http.StatusInternalServerError, "error creating patient", err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, output)
}

func (h *PatientHandler) ListPatients(c *gin.Context) {
	patients, err := h.listPatientsUC.Execute()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error listing patients", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, patients)
}

func (h *PatientHandler) GetPatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid patient ID", err)
		return
	}

	patient, err := h.getPatientUC.Execute(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "patient not found", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, patient)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid patient ID", err)
		return
	}

	var req UpdatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
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
		response.ErrorResponse(c, http.StatusInternalServerError, "error updating patient", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *PatientHandler) PatchPatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid patient ID", err)
		return
	}

	var req PatchPatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.patchPatientUC.Execute(application.PatchPatientInput{
		ID:        uint(id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Dni:       req.Dni,
		Phone:     req.Phone,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error patching patient", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}
