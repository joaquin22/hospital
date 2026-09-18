package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/doctors/application"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
)

type RegisterDoctorExecutor interface {
	Execute(input application.RegisterDoctorInput) (*application.RegisterDoctorOutput, error)
}

type UpdateDoctorExecutor interface {
	Execute(input application.UpdateDoctorInput) (*application.DoctorUserOutput, error)
}

type DoctorHandler struct {
	listDoctorUC     *application.ListDoctorUseCase
	getDoctorUC      *application.GetDoctorUseCase
	updateDoctorUC   UpdateDoctorExecutor
	patchDoctorUC    *application.PatchDoctorUseCase
	registerDoctorUC RegisterDoctorExecutor
}

func NewDoctorHandler(listDoctorUC *application.ListDoctorUseCase, getDoctorUC *application.GetDoctorUseCase, updateDoctorUC UpdateDoctorExecutor, patchDoctorUC *application.PatchDoctorUseCase, registerDoctorUC RegisterDoctorExecutor) *DoctorHandler {
	return &DoctorHandler{
		listDoctorUC:     listDoctorUC,
		getDoctorUC:      getDoctorUC,
		updateDoctorUC:   updateDoctorUC,
		patchDoctorUC:    patchDoctorUC,
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

func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID", err)
		return
	}

	output, err := h.getDoctorUC.Execute(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "Doctor not found", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID", err)
		return
	}

	var updateDoctorRequest UpdateDoctorRequest

	if err := c.ShouldBindJSON(&updateDoctorRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	output, err := h.updateDoctorUC.Execute(application.UpdateDoctorInput{
		ID:            uint(id),
		FirstName:     updateDoctorRequest.FirstName,
		LastName:      updateDoctorRequest.LastName,
		Email:         updateDoctorRequest.Email,
		Password:      updateDoctorRequest.Password,
		Role:          updateDoctorRequest.Role,
		Dni:           updateDoctorRequest.Dni,
		SpecialtyID:   updateDoctorRequest.SpecialtyID,
		LicenseNumber: updateDoctorRequest.LicenseNumber,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to update doctor", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *DoctorHandler) PatchDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid doctor ID", err)
		return
	}

	var req PatchDoctorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	output, err := h.patchDoctorUC.Execute(application.PatchDoctorInput{
		ID:            uint(id),
		SpecialtyID:   req.SpecialtyID,
		LicenseNumber: req.LicenseNumber,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to patch doctor", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}
