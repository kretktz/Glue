package service

import (
	"errors"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

type ITicketService interface {
	ListAllAvailableTickets() ([]entity.ITicket, error)

	CreateNewTicketPsql(ticket *entity.ITicket) (*entity.ITicket, error)
	ValidateTicketPsql(ticket *entity.ITicket) error
}

var ticketRepo repository.ITicketRepository

//TicketService creates a new service for ITicket
func TicketService(repository repository.ITicketRepository) ITicketService {
	ticketRepo = repository
	return &service{}
}


func (*service) ListAllAvailableTickets() ([]entity.ITicket, error) {
	return ticketRepo.ListAllAvailableTickets()
}

func (*service) CreateNewTicketPsql(ticket *entity.ITicket) (*entity.ITicket, error) {
	return ticketRepo.CreateNewTicketPsql(ticket)
}

func (*service) ValidateTicketPsql(ticket *entity.ITicket) error {
	if ticket == nil {
		err := errors.New("the ticket is not specified")
		return err
	}
	return nil
}