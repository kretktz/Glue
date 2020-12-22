package controller

import (
	"encoding/json"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/errors"
	"glue/glue-backend-golang/service"
	"net/http"
)

type controller struct{}

var (
	placeService service.PlaceService
)

//PlaceController interface to implement GetPlaces and AddPlace methods
type PlaceController interface {
	FireStoreGetPlaces(res http.ResponseWriter, req *http.Request)
	FireStoreAddPlace(res http.ResponseWriter, req *http.Request)
}

//NewPlaceController returns controller
func NewPlaceController(service service.PlaceService) PlaceController {
	placeService = service
	return &controller{}
}

// FirestoreGetPlaces lists all places
func (*controller) FireStoreGetPlaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	places, err := placeService.FireStoreFindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the places"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(places)
}

// FirestoreAddPlace adds a new place record to the Firestore DB
func (*controller) FireStoreAddPlace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var place entity.Place
	err := json.NewDecoder(req.Body).Decode(&place)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := placeService.FireStoreValidate(&place)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := placeService.FireStoreCreate(&place)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the place"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)

}
