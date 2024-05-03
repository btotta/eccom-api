package handler

import (
	"eccom-api/internal/domain/dtos"
	"eccom-api/internal/domain/entity"
	"eccom-api/internal/domain/repository"
	common_error "eccom-api/internal/pkg/erros"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAddressHandler struct {
	userAddressRepository repository.UserAddressRepository
	addressRepository     repository.AddressRepository
	userRepository        repository.UserRepository
}

func NewUserAddressHandler(userAddressRepository repository.UserAddressRepository,
	addressRepository repository.AddressRepository, userRepository repository.UserRepository) UserAddressHandler {

	return UserAddressHandler{
		userAddressRepository: userAddressRepository,
		addressRepository:     addressRepository,
		userRepository:        userRepository,
	}
}

// @Summary Create User Address
// @Description Create a new user address
// @Tags User
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserAddressDTO true "User object that needs to be created"
// @Success 201 {object} dtos.UserAddressResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 401 {object} common_error.ErrorResponse
// @Router /user/address [post]
func (h *UserAddressHandler) CreateUserAddress(c *gin.Context) {

	email := c.GetString("email")
	if email == "" {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "usuario não autenticado")
		return
	}

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "usuario não encontrado")
		return
	}

	var dto dtos.CreateUserAddressDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	state, err := h.addressRepository.FindStateByID(dto.StatedId)
	if err != nil && state.Status != entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "estado não encontrado")
		return
	}

	city, err := h.addressRepository.FindCityByID(dto.CityId)
	if err != nil && city.Status != entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "cidade não encontrada")
		return
	}

	neighborhood, err := h.addressRepository.FindNeighborhoodByID(dto.NeighborhoodId)
	if err != nil && neighborhood.Status != entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "bairro não encontrado")
		return
	}

	place, err := h.addressRepository.FindPlaceByID(dto.PlaceId)
	if err != nil && place.Status != entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "localidade não encontrada")
		return
	}

	userAddress := entity.UserAddress{
		StateID:        dto.StatedId,
		CityID:         dto.CityId,
		NeighborhoodID: dto.NeighborhoodId,
		PlaceID:        dto.PlaceId,
		Numero:         dto.Numero,
		Complemento:    dto.Complemento,
		UserID:         user.ID,
		Status:         entity.Active,
	}

	if _, err := h.userAddressRepository.CreateUserAddress(&userAddress); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.NewUserAddressResponseDTO(&userAddress))
}

// @Summary Get User Address
// @Description Get user address by id
// @Tags User
// @Accept json
// @Produce json
// @Param id query int true "User Address ID"
// @Success 200 {object} dtos.UserAddressResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 401 {object} common_error.ErrorResponse
// @Router /user/address [get]
func (h *UserAddressHandler) GetUserAddress(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "usuario não autenticado")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "usuario não encontrado")
		return
	}

	userAddress, err := h.userAddressRepository.FindUserAddressByIDAndUserID(uint(id), user.ID)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "endereço não encontrado")
		return
	}

	c.JSON(http.StatusOK, dtos.NewUserAddressResponseDTO(userAddress))
}

// @Summary Delete User Address
// @Description Delete user address by id
// @Tags User
// @Accept json
// @Produce json
// @Param id query int true "User Address ID"
// @Success 200 {object} dtos.UserAddressResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 401 {object} common_error.ErrorResponse
// @Router /user/address [delete]
func (h *UserAddressHandler) DeleteUserAddress(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "usuario não autenticado")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "id inválido")
		return
	}

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "usuario não encontrado")
		return
	}

	userAddress, err := h.userAddressRepository.FindUserAddressByIDAndUserID(uint(id), user.ID)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "endereço não encontrado")
		return
	}

	userAddress.Status = entity.Deleted

	if _, err := h.userAddressRepository.SaveUserAddress(userAddress); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewUserAddressResponseDTO(userAddress))
}

// @Summary Get User Address Page
// @Description Get user address by page
// @Tags User
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param size query int true "Page size"
// @Param sort query []string false "Sort" collectionFormat(multi)
// @Success 200 {object} dtos.UserAddressResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 401 {object} common_error.ErrorResponse
// @Router /user/address/paginated [get]
func (h *UserAddressHandler) GetUserAddressPage(c *gin.Context) {
	email := c.GetString("email")
	if email == "" {
		common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "usuario não autenticado")
		return
	}

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "usuario não encontrado")
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "page inválido")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("size", "15"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "limit inválido")
		return
	}

	orderBy, _ := c.GetQueryArray("sort")

	userAddresses, err := h.userAddressRepository.GetUserAddresses(user.ID, page, limit, orderBy)
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewPageUserAddressResponseDTO(userAddresses))
}
