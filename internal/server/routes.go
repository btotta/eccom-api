package server

import (
	"eccom-api/internal/domain/handler"
	"eccom-api/internal/domain/repository"
	"eccom-api/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	helloRepository repository.HelloRepository
	helloHandler    handler.HelloHandler
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	s.startRepositorys()
	s.startHandlers()

	hello := r.Group("")
	{
		hello.GET("/hello", helloHandler.Hello)
		hello.GET("/health", helloHandler.Health)
	}

	return r
}

func (s *Server) startRepositorys() {
	if helloRepository == nil {
		helloRepository = repository.NewHelloRepository(s.db.GetDB())
	}
}

func (s *Server) startHandlers() {
	if helloHandler == nil {
		helloHandler = handler.NewHelloHandler(helloRepository)
	}
}
