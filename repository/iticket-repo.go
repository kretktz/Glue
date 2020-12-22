package repository

import (
	"glue/glue-backend-golang/entity"
)
// ITicketRepository implements methods concerning tickets
type ITicketRepository interface {
	// Firestore methods
	FireStoreListAllAvailableTickets() ([]entity.ITicket, error)

	// PostgreSQL methods
	PsqlCreateNewTicket(ticket *entity.ITicket) (*entity.ITicket, error)
}
