package entity

// ISpace struct with the field details
type ISpace struct {
	Address                string   `json:"address"`
	Availability           string   `json:"availability"`
	Coordinates            string   `json:"coordinates"`
	Description            string   `json:"description"`
	ImageURLS              []string `json:"imageUrls"`
	Location               string   `json:"location"`
	Name                   string   `json:"name"`
	NumberOfVisitors       string   `json:"numberOfVisitors"`
	TelephoneNumber        string   `json:"telephoneNumber"`
	Tickets                []Ticket `json:"tickets"`
	TopImageURL            string   `json:"topImageUrl"`
	UID                    string   `json:"uid"`
	VisitorGreeting        string   `json:"visitorGreeting"`
	VisitorSlackMessage    string   `json:"visitorSlackMessage"`
	VisitorSlackWebhookURL string   `json:"visitorSlackWebhookUrl"`
	Website                string   `json:"website"`
}
