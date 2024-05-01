package handler

import (
	"eccom-api/internal/domain/dtos"
	"eccom-api/internal/domain/entity"
	"eccom-api/internal/domain/repository"
	common_error "eccom-api/internal/pkg/erros"
	"eccom-api/internal/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)
	RefreshTokenUser(c *gin.Context)
}

type userHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) UserHandler {
	return &userHandler{
		userRepository: userRepository,
	}
}

// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserDTO true "User object that needs to be created"
// @Success 201 {object} dtos.UserResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /user [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	var createUserDTO dtos.CreateUserDTO
	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := createUserDTO.ToEntity()
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := h.userRepository.GetUserByEmail(user.Email); err == nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "já existe um usuário com este e-mail")
		return
	}

	if err := h.userRepository.CreateUser(user); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.NewUserResponseDTO(user))
}

// @Summary Delete User
// @Description Delete a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} common_error.ErrorResponse
// @Router /user/{id} [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if id == 0 || err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	user, err := h.userRepository.GetUserByID(id)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "usuário não encontrado")
		return
	}

	user.Status = entity.Deleted
	if err := h.userRepository.SaveUser(user); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get User
// @Description Get a user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /user [get]
func (h *userHandler) GetUser(c *gin.Context) {

	email := c.GetString("email")
	if email == "" {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "usuário não autenticado")
		return
	}

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "usuário não encontrado")
		return
	}

	c.JSON(http.StatusOK, dtos.NewUserResponseDTO(user))
}

// @Summary Login User
// @Description Login a user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dtos.UserAuthDTO true "User credentials to login"
// @Success 200 {object} dtos.UserResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /user/login [post]
func (h *userHandler) LoginUser(c *gin.Context) {

	var userAuthDTO dtos.UserAuthDTO
	if err := c.ShouldBindJSON(&userAuthDTO); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userRepository.GetUserByEmail(userAuthDTO.Email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "usuário não encontrado")
		return
	}

	if err := user.ComparePassword(userAuthDTO.Password); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "credenciais inválidas")
		return
	}

	token, err := jwt.GenerateJwtToken(user.Email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refresh, err := jwt.GenerateRefreshToken(user.Email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewUserAuthResponseDTO(token, refresh))
}

// @Summary Refresh Token
// @Description Refresh a token
// @Tags User
// @Accept json
// @Produce json
// @Param refresh_token body dtos.UserAuthRefreshDTO true " Refresh token to refresh"
// @Success 200 {object} dtos.UserResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /user/refresh [post]
func (h *userHandler) RefreshTokenUser(c *gin.Context) {

	var tokens dtos.UserAuthRefreshDTO
	if err := c.ShouldBindJSON(&tokens); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "credenciais inválidas")
		return
	}

	token, err := jwt.RefreshToken(tokens.RefreshToken)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "credenciais inválidas")
		return
	}

	c.JSON(http.StatusOK, dtos.NewUserAuthResponseDTO(token, tokens.RefreshToken))
}

// @Summary Logout User
// @Description Logout a user
// @Tags User
// @Accept json
// @Produce json
// @Param refresh_token body dtos.UserAuthRefreshDTO true "Refresh token to logout"
// @Success 200
// @Failure 400 {object} common_error.ErrorResponse
// @Router /user/logout [post]
func (h *userHandler) LogoutUser(c *gin.Context) {
	var tokens dtos.UserAuthRefreshDTO
	if err := c.ShouldBindJSON(&tokens); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "credenciais inválidas")
		return
	}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "credenciais inválidas")
		return
	}

	jwt.InvalidateToken(tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
