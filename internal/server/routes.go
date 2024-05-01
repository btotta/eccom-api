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
	// Repository
	helloRepository repository.HelloRepository
	userRepository  repository.UserRepository

	// Handler
	helloHandler handler.HelloHandler
	userHandler  handler.UserHandler
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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

	user := r.Group("/user")
	{
		user.POST("", userHandler.CreateUser)
		user.DELETE("/:id", middleware.AuthMiddleware(), userHandler.DeleteUser)
		user.GET("", middleware.AuthMiddleware(), userHandler.GetUser)
		//user.PUT("", middleware.AuthMiddleware(), userHandler.UpdateUser)

		user.POST("/login", userHandler.LoginUser)
		user.POST("/logout", middleware.AuthMiddleware(), userHandler.LogoutUser)
		user.POST("/refresh", userHandler.RefreshTokenUser)
	}

	return r
}

func (s *Server) startRepositorys() {
	if helloRepository == nil {
		helloRepository = repository.NewHelloRepository(s.db.GetDB())
	}
	if userRepository == nil {
		userRepository = repository.NewUserRepository(s.db.GetDB())
	}
}

func (s *Server) startHandlers() {
	if helloHandler == nil {
		helloHandler = handler.NewHelloHandler(helloRepository)
	}
	if userHandler == nil {
		userHandler = handler.NewUserHandler(userRepository)
	}
}
