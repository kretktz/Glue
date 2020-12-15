package controller

import (
	"encoding/json"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/errors"
	"glue/glue-backend-golang/service"
	"net/http"
)

var (
	spaceService service.ISpaceService
	spaceID string
)

//ISpaceController interface to implement ListSpaces and GetSpaceByID method
type ISpaceController interface {
	ListSpaces(res http.ResponseWriter, req *http.Request)
	GetSpaceByID(res http.ResponseWriter, req *http.Request)
	CreateNewSpace(res http.ResponseWriter, req *http.Request)

	ListSpacesPsql(res http.ResponseWriter, req *http.Request)
	CreateNewSpacePsql(res http.ResponseWriter, req *http.Request)
}

//NewISpaceController returns controller
func NewISpaceController(service service.ISpaceService) ISpaceController {
	spaceService = service
	return &controller{}
}

// ListSpaces lists all spaces
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

func (*controller) ListSpacesPsql(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaces, err := spaceService.ListSpacesPsql()
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
	spaceIDs := req.URL.Query()["spaceID"]
	spaceID := spaceIDs[0]
	space, err := spaceService.GetSpaceByID(string(spaceID))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the requested space"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(space)
}

// CreatNewSpace adds a new space
func (*controller) CreateNewSpace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var space entity.ISpace
	err := json.NewDecoder(req.Body).Decode(&space)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := spaceService.ValidateSpace(&space)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := spaceService.CreateSpace(&space)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the space"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)

}

func (*controller) CreateNewSpacePsql(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var space entity.ISpace
	err := json.NewDecoder(req.Body).Decode(&space)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := spaceService.ValidateSpace(&space)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := spaceService.CreateNewSpacePsql(&space)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the space"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
