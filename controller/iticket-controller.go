package controller

import (
	"encoding/json"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/errors"
	"glue/glue-backend-golang/service"
	"net/http"
)

var ticketService service.ITicketService

//ISpaceController interface to implement ListSpaces and GetSpaceByID method
type ITicketController interface {
	FireStoreListAllAvailableTickets(res http.ResponseWriter, req *http.Request)

	PsqlCreateNewTicket(res http.ResponseWriter, req *http.Request)
}

//NewISpaceController returns controller
func NewITicketController(service service.ITicketService) ITicketController {
	ticketService = service
	return &controller{}
}

// ListSpaces lists spaces
func (*controller) FireStoreListAllAvailableTickets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	tickets, err := ticketService.FireStoreListAllAvailableTickets()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the tickets"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(tickets)
}

func (*controller) PsqlCreateNewTicket(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var ticket entity.ITicket
	err := json.NewDecoder(req.Body).Decode(&ticket)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := ticketService.ValidateTicketPsql(&ticket)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := ticketService.PsqlCreateNewTicket(&ticket)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the space"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
