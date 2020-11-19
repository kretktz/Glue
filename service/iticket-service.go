package service

import (
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

type ITicketService interface {
	ListAllAvailableTickets() ([]entity.ITicket, error)
}

var ticketRepo repository.ITicketRepository

//TicketService creates a new service for ISpace
func TicketService(repository repository.ITicketRepository) ITicketService {
	ticketRepo = repository
	return &service{}
}


func (*service) ListAllAvailableTickets() ([]entity.ITicket, error) {
	return ticketRepo.ListAllAvailableTickets()
}