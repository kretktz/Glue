package entity

//ITenant struct containing tenants' data
type ITenant struct {
	ExpiryDate             string `json:"expiry_date"`
	Hidden                 bool   `json:"hidden"`
	UID                    string `json:"uid"`
	UserID                 string `json:"user_id"`
	VisitorGreeting        string `json:"visitor_greeting,omitempty"`
	VisitorSlackMessage    string `json:"visitor_slack_message,omitempty"`
	VisitorSlackWebHookURL string `json:"visitor_slack_web_hook_url,omitempty"`
}
