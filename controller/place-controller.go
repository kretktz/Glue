package controller

import (
	"encoding/json"
	entity "glue/glue-backend-golang/entity"
	errors "glue/glue-backend-golang/errors"
	service "glue/glue-backend-golang/service"
	"net/http"
)

type controller struct{}

var (
	placeService service.PlaceService
)

//PlaceController interface to implement GetPlaces and AddPlace methods
type PlaceController interface {
	GetPlaces(res http.ResponseWriter, req *http.Request)
	AddPlace(res http.ResponseWriter, req *http.Request)
}

//NewPlaceController returns controller
func NewPlaceController(service service.PlaceService) PlaceController {
	placeService = service
	return &controller{}
}

// GetPlaces gets places
func (*controller) GetPlaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	places, err := placeService.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the places"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(places)
}

// AddPlace adds a place
func (*controller) AddPlace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var place entity.IPlace
	err := json.NewDecoder(req.Body).Decode(&place)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := placeService.Validate(&place)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := placeService.Create(&place)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the place"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)

}
