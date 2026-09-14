package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/doctors/application"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
)

type RegisterDoctorExecutor interface {
	Execute(input application.RegisterDoctorInput) (*application.RegisterDoctorOutput, error)
}

type DoctorHandler struct {
	listDoctorUC     *application.ListDoctorUseCase
	registerDoctorUC RegisterDoctorExecutor
}

func NewDoctorHandler(listDoctorUC *application.ListDoctorUseCase, registerDoctorUC RegisterDoctorExecutor) *DoctorHandler {
	return &DoctorHandler{
		listDoctorUC:     listDoctorUC,
		registerDoctorUC: registerDoctorUC,
	}
}

func (h *DoctorHandler) RegisterDoctor(c *gin.Context) {

	var createDoctorRequest CreateDoctorRequest

	if err := c.ShouldBindJSON(&createDoctorRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	output, err := h.registerDoctorUC.Execute(application.RegisterDoctorInput{
		FirstName:     createDoctorRequest.FirstName,
		LastName:      createDoctorRequest.LastName,
		Email:         createDoctorRequest.Email,
		Password:      createDoctorRequest.Password,
		Role:          createDoctorRequest.Role,
		Dni:           createDoctorRequest.Dni,
		SpecialtyID:   createDoctorRequest.SpecialtyID,
		LicenseNumber: createDoctorRequest.LicenseNumber,
	})

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to register doctor", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *DoctorHandler) ListDoctors(c *gin.Context) {
	output, err := h.listDoctorUC.Execute()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to list doctors", err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, output)
}
