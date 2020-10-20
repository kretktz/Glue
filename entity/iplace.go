package entity

// IPlace struct
type IPlace struct {
	ConfirmPageTitle string `json:"confirmPageTitle"`
	PhoneNumber      string `json:"phoneNumber"`
	SlackSentMessage string `json:"slackSentMessage"`
	SlackWebHookURL  string `json:"slackWebHookURL"`
	VisitPlaceName   string `json:"visitPlaceName"`
}
