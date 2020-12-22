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

//ISpaceController interface to implement ISpace related methods
type ISpaceController interface {
	FireStoreListSpaces(res http.ResponseWriter, req *http.Request)
	FireStoreGetSpaceByID(res http.ResponseWriter, req *http.Request)
	FireStoreCreateNewSpace(res http.ResponseWriter, req *http.Request)

	PsqlListSpaces(res http.ResponseWriter, req *http.Request)
	PsqlCreateNewSpace(res http.ResponseWriter, req *http.Request)
	PsqlGetSpaceByID(res http.ResponseWriter, req *http.Request)
	PsqlListSpacesWithTickets(res http.ResponseWriter, req *http.Request)

}

//NewISpaceController returns controller
func NewISpaceController(service service.ISpaceService) ISpaceController {
	spaceService = service
	return &controller{}
}

// ListSpaces lists all spaces
func (*controller) FireStoreListSpaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaces, err := spaceService.FireStoreListSpaces()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the spaces"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(spaces)
}

func (*controller) PsqlListSpaces(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaces, err := spaceService.PsqlListSpaces()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the spaces"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(spaces)
}

func (*controller) PsqlListSpacesWithTickets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaces, tickets, err := spaceService.PsqlListSpacesWithTickets()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the spaces"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(spaces)
	json.NewEncoder(res).Encode(tickets)
}

// GetSpaceByID gets a particular space as specified by provided UID
func (*controller) FireStoreGetSpaceByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaceIDs := req.URL.Query()["spaceID"]
	spaceID := spaceIDs[0]
	space, err := spaceService.FireStoreGetSpaceByID(string(spaceID))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the requested space"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(space)
}

func (*controller) PsqlGetSpaceByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	spaceIDs := req.URL.Query()["spaceID"]
	spaceID := spaceIDs[0]
	space, err := spaceService.PsqlGetSpaceByID(string(spaceID))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the requested space"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(space)
}

// CreatNewSpace adds a new space
func (*controller) FireStoreCreateNewSpace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var space entity.ISpace
	err := json.NewDecoder(req.Body).Decode(&space)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := spaceService.FireStoreValidateSpace(&space)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := spaceService.FireStoreCreateSpace(&space)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the space"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)

}

func (*controller) PsqlCreateNewSpace(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var space entity.ISpace
	err := json.NewDecoder(req.Body).Decode(&space)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := spaceService.FireStoreValidateSpace(&space)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := spaceService.PsqlCreateNewSpace(&space)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the space"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
