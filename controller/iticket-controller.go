package controller

import (
	"encoding/json"
	"glue/glue-backend-golang/errors"
	"glue/glue-backend-golang/service"
	"net/http"
)

var ticketService service.ITicketService

//ISpaceController interface to implement ListSpaces and GetSpaceByID method
type ITicketController interface {
	ListAllAvailableTickets(res http.ResponseWriter, req *http.Request)
}

//NewISpaceController returns controller
func NewITicketController(service service.ITicketService) ITicketController {
	ticketService = service
	return &controller{}
}

// ListSpaces lists spaces
func (*controller) ListAllAvailableTickets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	tickets, err := ticketService.ListAllAvailableTickets()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error fetching the tickets"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(tickets)
}
