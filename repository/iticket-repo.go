package repository

import (
	"glue/glue-backend-golang/entity"
)

type ITicketRepository interface {
	FireStoreListAllAvailableTickets() ([]entity.ITicket, error)

	PsqlCreateNewTicket(ticket *entity.ITicket) (*entity.ITicket, error)
}
