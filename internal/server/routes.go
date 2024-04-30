package server

import (
	"eccom-api/internal/domain/handler"
	"eccom-api/internal/domain/repository"
	"eccom-api/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "eccom-api/docs"
)

var (
	helloRepository repository.HelloRepository
	helloHandler    handler.HelloHandler
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.OPTIONS("/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	s.startRepositorys()
	s.startHandlers()

	hello := r.Group("")
	{
		hello.GET("/", helloHandler.Hello)
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
