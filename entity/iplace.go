package entity

// IPlace struct
type IPlace struct {
	ConfirmPageTitle string `json:"confirm_page_title"`
	PhoneNumber      string `json:"phone_number"`
	SlackSentMessage string `json:"slack_sent_message"`
	SlackWebHookURL  string `json:"slack_web_hook_URL"`
	VisitPlaceName   string `json:"visit_place_name"`
}
