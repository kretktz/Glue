package entity

//Ticket struct containing the details of a ticket
type Ticket struct {
	PlaceName          string `json:"placeName"`
	TicketType         string `json:"ticketType"`
	NumberTicketsAvail int64  `json:"numberTicketsAvail"`
}
