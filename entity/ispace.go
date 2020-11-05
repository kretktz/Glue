package entity

// ISpace struct with the field details
type ISpace struct {
	Address                string   `json:"address"`
	Availability           string   `json:"availability"`
	Coordinates            string   `json:"coordinates"`
	Description            string   `json:"description"`
	ImageURLS              string   `json:"image_URLs"`
	Location               string   `json:"location"`
	Name                   string   `json:"name"`
	NumberOfVisitors       string   `json:"number_of_visitors"`
	TelephoneNumber        string   `json:"telephone_number"`
	Tickets                ITicket `json:"tickets"`
	TopImageURL            string   `json:"top_image_URL"`
	UID                    string   `json:"uid"`
	VisitorGreeting        string   `json:"visitor_greeting"`
	VisitorSlackMessage    string   `json:"visitor_slack_message"`
	VisitorSlackWebhookURL string   `json:"visitor_slack_webhook_URL"`
	Website                string   `json:"website"`
}
