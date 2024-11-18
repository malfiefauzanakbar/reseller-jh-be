package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

// type BaseResp struct {
// 	Status     string      `json:"status"`
// 	Message    string      `json:"message,omitempty"`
// 	Data       interface{} `json:"data,omitempty"`
// 	Pagination *Pagination `json:"pagination,omitempty"`
// 	Errors     interface{} `json:"errors,omitempty"`
// }

type SuccessResponse struct {
	Success    bool        `json:"success"`
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// ErrorResponse defines the structure of an error response
type ErrorResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

// RespondSuccess sends a standardized success response
func RespondSuccess(c *gin.Context, message string, data interface{}, pagination *Pagination) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success:    true,
		Status:     http.StatusOK,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// RespondError sends a standardized error response
func RespondError(c *gin.Context, httpStatus int, message string, err interface{}) {
	c.JSON(httpStatus, ErrorResponse{
		Success: false,
		Status:  httpStatus,
		Message: message,
		Error:   err,
	})
}
