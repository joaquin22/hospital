package http

type CreateSpecialityRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateSpecialityRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PatchSpecialityRequest struct {
	ID          uint    `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
