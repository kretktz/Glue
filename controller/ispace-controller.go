package controller

import (
	"encoding/json"
	"glue/glue-backend-golang/errors"
	"glue/glue-backend-golang/service"
	"net/http"
)

var (
	spaceService service.ISpaceService
)

//ISpaceController interface to implement ListSpaces and GetSpaceByID method
type ISpaceController interface {
	ListSpaces(res http.ResponseWriter, req *http.Request)
	GetSpaceByID(res http.ResponseWriter, req *http.Request)
}

//NewISpaceController returns controller
func NewISpaceController(service service.ISpaceService) ISpaceController {
	spaceService = service
	return &controller{}
}

// ListSpaces lists spaces
func (*controller) ListSpaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaces, err := spaceService.ListSpaces()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the spaces"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(spaces)
}

// GetSpaceByID gets a particular space as specified by provided UID
func (*controller) GetSpaceByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaces, err := spaceService.GetSpaceByID()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the requested space"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(spaces)
}
