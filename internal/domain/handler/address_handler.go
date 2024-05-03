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

type AddressHandler struct {
	addressRepository repository.AddressRepository
}

func NewAddressHandler(addressRepository repository.AddressRepository) AddressHandler {
	return AddressHandler{
		addressRepository: addressRepository,
	}
}

// @Summary Create a state
// @Description Create a state
// @Tags Address
// @Accept json
// @Produce json
// @Param state body dtos.CreateStateDTO true "State"
// @Success 201 {object} dtos.StateResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/state [post]
func (h *AddressHandler) CreateState(c *gin.Context) {

	var state dtos.CreateStateDTO

	if err := c.ShouldBindJSON(&state); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	stateEntity := state.ToState()

	if state, _ := h.addressRepository.FindStateByUF(stateEntity.UF); state != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "Estado já cadastrado")
		return
	}

	if _, err := h.addressRepository.SaveState(stateEntity); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.NewStateResponseDTO(stateEntity))

}

// @Summary Create a city
// @Description Create a city
// @Tags Address
// @Accept json
// @Produce json
// @Param city body dtos.CreateCityDTO true "City"
// @Success 201 {object} dtos.CityResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/city [post]
func (h *AddressHandler) CreateCity(c *gin.Context) {

	var city dtos.CreateCityDTO

	if err := c.ShouldBindJSON(&city); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cityEntity := city.ToCity()

	state, err := h.addressRepository.FindStateByID(city.StateID)
	if err != nil || state.Status == entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "Estado não encontrado")
		return
	}

	cityEntity.StateID = state.ID

	if city, _ := h.addressRepository.FindCityByName(cityEntity.Name); city != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "Cidade já cadastrada")
		return
	}

	if _, err := h.addressRepository.SaveCity(cityEntity); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.NewCityResponseDTO(cityEntity))

}

// @Summary Create a neighborhood
// @Description Create a neighborhood
// @Tags Address
// @Accept json
// @Produce json
// @Param neighborhood body dtos.CreateNeighborhoodDTO true "Neighborhood"
// @Success 201 {object} dtos.NeighborhoodResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/neighborhood [post]//
func (h *AddressHandler) CreateNeighborhood(c *gin.Context) {

	var neighborhood dtos.CreateNeighborhoodDTO

	if err := c.ShouldBindJSON(&neighborhood); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	city, err := h.addressRepository.FindCityByID(neighborhood.CityID)
	if err != nil || city.Status == entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "Cidade não encontrada")
		return
	}

	neighborhoodEntity := neighborhood.ToNeighborhood()

	neighborhoodEntity.CityID = city.ID

	if _, err := h.addressRepository.SaveNeighborhood(neighborhoodEntity); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.NewNeighborhoodResponseDTO(neighborhoodEntity))

}

// @Summary Create a place
// @Description Create a place
// @Tags Address
// @Accept json
// @Produce json
// @Param place body dtos.CreatePlaceDTO true "Place"
// @Success 201 {object} dtos.PlaceResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/place [post]
func (h *AddressHandler) CreatePlace(c *gin.Context) {

	var place dtos.CreatePlaceDTO

	if err := c.ShouldBindJSON(&place); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	neighborhood, err := h.addressRepository.FindNeighborhoodByID(place.NeighborhoodID)
	if err != nil || neighborhood.Status == entity.Deleted {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "Bairro não encontrado")
		return
	}

	placeEntity := place.ToPlace()

	placeEntity.NeighborhoodId = neighborhood.ID

	if _, err := h.addressRepository.SavePlace(placeEntity); err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dtos.NewPlaceResponseDTO(placeEntity))

}

// @Summary Get a state
// @Description Get a state
// @Tags Address
// @Accept json
// @Produce json
// @Param id path int true "State ID"
// @Success 200 {object} dtos.StateResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 404 {object} common_error.ErrorResponse
// @Router /address/state/{id} [get]
func (h *AddressHandler) GetState(c *gin.Context) {

	stateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	state, err := h.addressRepository.FindStateByID(uint(stateID))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "Estado não encontrado")
		return
	}

	c.JSON(http.StatusOK, dtos.NewStateResponseDTO(state))

}

// @Summary Get a city
// @Description Get a city
// @Tags Address
// @Accept json
// @Produce json
// @Param id path int true "City ID"
// @Success 200 {object} dtos.CityResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 404 {object} common_error.ErrorResponse
// @Router /address/city/{id} [get]
func (h *AddressHandler) GetCity(c *gin.Context) {

	cityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	city, err := h.addressRepository.FindCityByID(uint(cityID))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "Cidade não encontrada")
		return
	}

	c.JSON(http.StatusOK, dtos.NewCityResponseDTO(city))
}

// @Summary Get a neighborhood
// @Description Get a neighborhood
// @Tags Address
// @Accept json
// @Produce json
// @Param id path int true "Neighborhood ID"
// @Success 200 {object} dtos.NeighborhoodResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 404 {object} common_error.ErrorResponse
// @Router /address/neighborhood/{id} [get]
func (h *AddressHandler) GetNeighborhood(c *gin.Context) {

	neighborhoodID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	neighborhood, err := h.addressRepository.FindNeighborhoodByID(uint(neighborhoodID))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "Bairro não encontrado")
		return
	}

	c.JSON(http.StatusOK, dtos.NewNeighborhoodResponseDTO(neighborhood))
}

// @Summary Get a place
// @Description Get a place
// @Tags Address
// @Accept json
// @Produce json
// @Param id path int true "Place ID"
// @Success 200 {object} dtos.PlaceResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Failure 404 {object} common_error.ErrorResponse
// @Router /address/place/{id} [get]
func (h *AddressHandler) GetPlace(c *gin.Context) {

	placeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	place, err := h.addressRepository.FindPlaceByID(uint(placeID))
	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusNotFound, "Localidade não encontrada")
		return
	}

	c.JSON(http.StatusOK, dtos.NewPlaceResponseDTO(place))
}

// @Summary Get a page of states
// @Description Get a page of states
// @Tags Address
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query []string false "Sort" collectionFormat(multi)
// @Param uf query string false "UF"
// @Param name query string false "Name"
// @Success 200 {object} dtos.PageStateResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/state/paginated [get]
func (h *AddressHandler) GetStatePage(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "15"))
	orderBy, _ := c.GetQueryArray("sort")

	uf, _ := c.GetQuery("uf")
	name, _ := c.GetQuery("name")

	states, err := h.addressRepository.FindAllStates(page, size, orderBy, uf, name)

	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewPageStateResponseDTO(states))
}

// @Summary Get a page of cities
// @Description Get a page of cities
// @Tags Address
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query []string false "Sort" collectionFormat(multi)
// @Param state_id query int false "State ID"
// @Param name query string false "Name"
// @Success 200 {object} dtos.PageCityResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/city/paginated [get]
func (h *AddressHandler) GetCityPage(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "15"))
	orderBy, _ := c.GetQueryArray("sort")

	stateID, _ := strconv.Atoi(c.DefaultQuery("state_id", "0"))
	name, _ := c.GetQuery("name")

	cities, err := h.addressRepository.FindAllCities(page, size, orderBy, uint(stateID), name)

	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewPageCityResponseDTO(cities))
}

// @Summary Get a page of neighborhoods
// @Description Get a page of neighborhoods
// @Tags Address
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query []string false "Sort" collectionFormat(multi)
// @Param city_id query int false "City ID"
// @Param name query string false "Name"
// @Success 200 {object} dtos.PageNeighborhoodResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/neighborhood/paginated [get]
func (h *AddressHandler) GetNeighborhoodPage(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "15"))
	orderBy, _ := c.GetQueryArray("sort")

	cityID, _ := strconv.Atoi(c.DefaultQuery("city_id", "0"))
	name, _ := c.GetQuery("name")

	neighborhoods, err := h.addressRepository.FindAllNeighborhoods(page, size, orderBy, uint(cityID), name)

	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewPageNeighborhoodResponseDTO(neighborhoods))
}

// @Summary Get a page of places
// @Description Get a page of places
// @Tags Address
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query []string false "Sort" collectionFormat(multi)
// @Param neighborhood_id query int false "Neighborhood ID"
// @Param name query string false "Name"
// @Success 200 {object} dtos.PagePlaceResponseDTO
// @Failure 400 {object} common_error.ErrorResponse
// @Router /address/place/paginated [get]
func (h *AddressHandler) GetPlacePage(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "15"))
	orderBy, _ := c.GetQueryArray("sort")

	neighborhoodID, _ := strconv.Atoi(c.DefaultQuery("neighborhood_id", "0"))
	name, _ := c.GetQuery("name")

	places, err := h.addressRepository.FindAllPlaces(page, size, orderBy, uint(neighborhoodID), name)

	if err != nil {
		common_error.DefaultErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.NewPagePlaceResponseDTO(places))
}
