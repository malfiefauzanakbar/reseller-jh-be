package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse defines the structure of a successful response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse defines the structure of an error response
type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

// RespondSuccess sends a standardized success response
func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// RespondError sends a standardized error response
func RespondError(c *gin.Context, httpStatus int, message string, err interface{}) {
	c.JSON(httpStatus, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
