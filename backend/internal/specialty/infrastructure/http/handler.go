package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaquin22/hospital-api/internal/shared/infrastructure/response"
	"github.com/joaquin22/hospital-api/internal/specialty/application"
)

type SpecialityHandler struct {
	listSpecialityUC   *application.ListSpecialitiesUseCase
	getSpecialityUC    *application.GetSpecialityUseCase
	createSpecialityUC *application.CreateSpecilityUseCase
	updateSpecialityUC *application.UpdateSpecialityUseCase
	patchSpecialityUC  *application.PatchSpecialityUseCase
}

func NewSpecialityHandler(
	createSpecialityUC *application.CreateSpecilityUseCase,
	updateSpecialityUC *application.UpdateSpecialityUseCase,
	patchSpecialityUC *application.PatchSpecialityUseCase,
	listSpecialityUC *application.ListSpecialitiesUseCase,
	getSpecialityUC *application.GetSpecialityUseCase,
) *SpecialityHandler {
	return &SpecialityHandler{
		createSpecialityUC: createSpecialityUC,
		updateSpecialityUC: updateSpecialityUC,
		patchSpecialityUC:  patchSpecialityUC,
		listSpecialityUC:   listSpecialityUC,
		getSpecialityUC:    getSpecialityUC,
	}
}

func (h *SpecialityHandler) CreateSpeciality(c *gin.Context) {

	var req CreateSpecialityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.createSpecialityUC.Execute(application.CreateSpecialityInput{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error creating speciality", err)
		return
	}
	response.SuccessResponse(c, http.StatusCreated, output)
}

func (h *SpecialityHandler) ListSpecialities(c *gin.Context) {
	output, err := h.listSpecialityUC.Execute()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error listing specialities", err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *SpecialityHandler) GetSpeciality(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	output, err := h.getSpecialityUC.Execute(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, "speciality not found", err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *SpecialityHandler) UpdateSpeciality(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	var req UpdateSpecialityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.updateSpecialityUC.Execute(application.UpdateSpecialityInput{
		ID:          uint(id),
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "error updating speciality", err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, output)
}

func (h *SpecialityHandler) PatchSpeciality(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	var req PatchSpecialityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.patchSpecialityUC.Execute(application.PatchSpecialityInput{
		ID:          uint(id),
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "error patching speciality", err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, output)
}
