package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/doctors/application"
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
	userDomain "github.com/joaquin22/hospital-api/internal/users/domain"
)

type RegisterDoctorExecutor interface {
	Execute(input application.RegisterDoctorInput) (*application.RegisterDoctorOutput, error)
}

type UpdateDoctorExecutor interface {
	Execute(input application.UpdateDoctorInput) (*application.DoctorUserOutput, error)
}

type PatchDoctorExecutor interface {
	Execute(input application.PatchDoctorInput) (*application.DoctorUserOutput, error)
}

type DoctorHandler struct {
	listDoctorUC     *application.ListDoctorUseCase
	getDoctorUC      *application.GetDoctorUseCase
	registerDoctorUC RegisterDoctorExecutor
	updateDoctorUC   UpdateDoctorExecutor
	patchDoctorUC    PatchDoctorExecutor
}

func NewDoctorHandler(listDoctorUC *application.ListDoctorUseCase, getDoctorUC *application.GetDoctorUseCase, updateDoctorUC UpdateDoctorExecutor, patchDoctorUC PatchDoctorExecutor, registerDoctorUC RegisterDoctorExecutor) *DoctorHandler {
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
		response.ErrorResponse(c, statusForRegisterError(err), "Failed to register doctor", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

// statusForRegisterError distingue los errores de validación (400) de los
// fallos de la verificación externa del DNI (502 Bad Gateway) para que el
// cliente no los confunda con un error del servidor.
func statusForRegisterError(err error) int {
	switch {
	case errors.Is(err, userDomain.ErrDniNotFound):
		return http.StatusBadRequest
	case errors.Is(err, userDomain.ErrDniVerificationFailed):
		return http.StatusBadGateway
	case errors.Is(err, userDomain.ErrWeakPassword):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// statusForWriteDoctorError separa los errores de validación y de dominio (400)
// de los fallos inesperados (500). La escritura de un doctor delega en el caso
// de uso de usuarios, así que aquí también puede aparecer un ErrWeakPassword.
func statusForWriteDoctorError(err error) int {
	switch {
	case errors.Is(err, userDomain.ErrWeakPassword),
		errors.Is(err, userDomain.ErrInvalidEmail),
		errors.Is(err, userDomain.ErrInvalidDni),
		errors.Is(err, userDomain.ErrInvalidName),
		errors.Is(err, userDomain.ErrInvalidRole),
		errors.Is(err, userDomain.ErrEmailAlreadyExists),
		errors.Is(err, userDomain.ErrDniAlreadyExists):
		return http.StatusBadRequest
	case errors.Is(err, userDomain.ErrUserNotFound), errors.Is(err, domain.ErrDoctorNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
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
		Active:        updateDoctorRequest.Active,
	})
	if err != nil {
		response.ErrorResponse(c, statusForWriteDoctorError(err), "Failed to update doctor", err)
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
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Password:      req.Password,
		Role:          req.Role,
		Dni:           req.Dni,
		SpecialtyID:   req.SpecialtyID,
		LicenseNumber: req.LicenseNumber,
		Active:        req.Active,
	})
	if err != nil {
		response.ErrorResponse(c, statusForWriteDoctorError(err), "Failed to patch doctor", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}
