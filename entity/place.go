package entity

// Place struct with details of a single space
type Place struct {
	ConfirmPageTitle string `json:"confirmPageTitle"`
	PhoneNumber      string `json:"phoneNumber"`
	VisitPlaceName   string `json:"visitPlaceName"`
	SlackSentMessage string `json:"slackSentMessage"`
	SlackWebHookURL  string `json:"slackWebHookURL"`
}
