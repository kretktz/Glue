package entity

//IVoucherJob struct
type IVoucherJob struct {
	PaidTicketID string `json:"paid_ticket_id"`
	UID string `json:"uid,omitempty"`
	UserID string `json:"user_ID"`
	SpaceID string   `json:"space_ID"`
	Days int64 `json:"days"`
}
