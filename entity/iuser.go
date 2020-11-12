package entity

//IUser structure with user details
type IUser struct {
	Email                  string `json:"email"`
	Name                   string `json:"name"`
	Uid                    string `json:"uid"`
	Hidden                 bool   `json:"hidden,omitempty"`
	VisitorGreeting        string `json:"visitor_greeting,omitempty"`
	VisitorSlackMessage    string `json:"visitor_slack_message,omitempty"`
	VisitorSlackWebHookURL string `json:"visitor_slack_web_hook_url,omitempty"`
}
