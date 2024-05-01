package common_error

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Error response DTO
type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
}

func newErrorResponse(c *gin.Context, status int, message string) *ErrorResponse {
	return &ErrorResponse{
		Timestamp: time.Now(),
		Status:    status,
		Error:     http.StatusText(status),
		Message:   message,
		Path:      c.Request.URL.Path,
	}
}

func DefaultErrorResponse(c *gin.Context, status int, message string) {
	errorHandler := newErrorResponse(c, status, "Ocorreu um erro: "+message)
	c.JSON(errorHandler.Status, errorHandler)
}
