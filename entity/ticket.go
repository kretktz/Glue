package entity

//Ticket struct containing the details of a ticket
type Ticket struct {
	TicketType         string `json:"ticketType"`
	NumberTicketsAvail int    `json:"numberTicketsAvail"`
}
