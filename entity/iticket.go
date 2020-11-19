package entity

// ITicket struct
type ITicket struct {
	Availability int64    `json:"availability"`
	Colour       string   `json:"colour"`
	Description  string   `json:"description"`
	Name         string   `json:"name"`
	Period       int64    `json:"period"`
	Price        int64    `json:"price"`
	SpaceID      string   `json:"space_ID"`
	Space        []ISpace `json:"space,omitempty"`
	UID          string   `json:"uid"`
}
