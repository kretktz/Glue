package entity

//IPurchasedTicket struct containing data for a purchased ticket
type IPurchasedTicket struct {
	ExpiryDate    string    `json:"expiry_date"`
	PaymentMethod string    `json:"payment_method"`
	PurchasedDate string    `json:"purchased_date"`
	Receipt       string    `json:"receipt"`
	TicketID      string    `json:"ticket_id"`
	Ticket        []ITicket `json:"ticket"`
	UID           string    `json:"uid"`
	UserID        string    `json:"user_id"`
}
