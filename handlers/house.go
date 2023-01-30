package handlers

import (
	"encoding/json"
	housesdto "housybe/dto/house"
	dto "housybe/dto/result"
	"housybe/models"
	"housybe/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type handlerHouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerHouse {
	return &handlerHouse{HouseRepository}
}

func (h *handlerHouse) FindHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	houses, err := h.HouseRepository.FindHouses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: houses}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) AddHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(housesdto.AddHouseRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house := models.House{
		Name:      request.Name,
		CityId:    request.CityId,
		Address:   request.Address,
		Price:     request.Price,
		TypeRent:  request.TypeRent,
		Ameneties: request.Ameneties,
		BedRoom:   request.BedRoom,
		BathRoom:  request.BathRoom,
		Image:     request.Image,
	}

	data, err := h.HouseRepository.AddHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseHouse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseHouse(u models.House) models.HouseResponse {
	return models.HouseResponse{
		Name:      u.Name,
		CityId:    u.CityId,
		Address:   u.Address,
		Price:     u.Price,
		TypeRent:  u.TypeRent,
		Ameneties: u.Ameneties,
		BedRoom:   u.BedRoom,
		BathRoom:  u.BathRoom,
		Image:     u.Image,
	}
}
