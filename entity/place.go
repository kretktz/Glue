package entity

// Place struct with details of a single place
type Place struct {
	PlaceName     string   `json:"placeName"`
	PlaceLocation string   `json:"placeLocation"`
	PhoneNumber   string   `json:"phoneNumber"`
	NumTickets    []Ticket `json:"numTickets"`
}
