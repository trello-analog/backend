package entity

import "github.com/trello-analog/backend/customerrors"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool                   `json:"success"`
	Error   *customerrors.APIError `json:"error"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(error *customerrors.APIError) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Error:   error,
	}
}
