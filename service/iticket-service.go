package service

import (
	"errors"
	"glue/glue-backend-golang/entity"
	"glue/glue-backend-golang/repository"
)

type ITicketService interface {
	FireStoreListAllAvailableTickets() ([]entity.ITicket, error)

	PsqlCreateNewTicket(ticket *entity.ITicket) (*entity.ITicket, error)
	ValidateTicketPsql(ticket *entity.ITicket) error
}

var ticketRepo repository.ITicketRepository

//TicketService creates a new service for ITicket
func TicketService(repository repository.ITicketRepository) ITicketService {
	ticketRepo = repository
	return &service{}
}


func (*service) FireStoreListAllAvailableTickets() ([]entity.ITicket, error) {
	return ticketRepo.FireStoreListAllAvailableTickets()
}

func (*service) PsqlCreateNewTicket(ticket *entity.ITicket) (*entity.ITicket, error) {
	return ticketRepo.PsqlCreateNewTicket(ticket)
}

func (*service) ValidateTicketPsql(ticket *entity.ITicket) error {
	if ticket == nil {
		err := errors.New("the ticket is not specified")
		return err
	}
	return nil
}