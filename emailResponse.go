package postmark

import "time"

// EmailResponse is the response from the Postmark API to an Email request
type EmailResponse struct {
	// The "to" address of the Email
	To string `json:"To"`
	// Postmark Message ID for created Email
	MessageID string `json:"MessageID"`

	// Time at which the email send request was submitted (RFC3339)
	SubmittedAt time.Time `json:"SubmittedAt"`
}
