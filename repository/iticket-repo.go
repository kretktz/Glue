package repository

import (
	"glue/glue-backend-golang/entity"
)

type ITicketRepository interface {
	ListAllAvailableTickets() ([]entity.ITicket, error)

	CreateNewTicketPsql(ticket *entity.ITicket) (*entity.ITicket, error)
}
