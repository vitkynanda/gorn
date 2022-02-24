package dtos

type ValidationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
