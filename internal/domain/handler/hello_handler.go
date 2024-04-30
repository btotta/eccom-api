package handler

import (
	"eccom-api/internal/domain/dtos"
	"eccom-api/internal/domain/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HelloHandler interface {
	Hello(c *gin.Context)
	Health(c *gin.Context)
}

type helloHandler struct {
	helloRepository repository.HelloRepository
}

func NewHelloHandler(helloRepository repository.HelloRepository) HelloHandler {
	return &helloHandler{
		helloRepository: helloRepository,
	}
}

// @Summary Hello
// @Description Hello
// @Tags Hello
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello World"
// @Router / [get]
func (h *helloHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// @Summary Health
// @Description Health
// @Tags Hello
// @Accept json
// @Produce json
// @Success 200 {object} dtos.Hello
// @Router /health [get]
func (h *helloHandler) Health(c *gin.Context) {
	if err := h.helloRepository.Health(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.Hello{
		Message:   "Hello World",
		Status:    "UP",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
