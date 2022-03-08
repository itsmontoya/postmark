package postmark

import "github.com/hatchify/errors"

const (
	// ErrEmptyToAddress is returned when an "to" address is empty
	ErrEmptyToAddress = errors.Error("invalid \"to\" address, cannot be empty")
	// ErrEmptyFromAddress is returned when an "from" address is empty
	ErrEmptyFromAddress = errors.Error("invalid \"from\" address, cannot be empty")
	// ErrEmptyBody is returned when an HTML body and text body are empty
	ErrEmptyBody = errors.Error("invalid body, both HTML body and text body cannot be empty")
)

// MakeEmail will initialize an Email value
func MakeEmail(from, to, subject, htmlBody string) (e Email) {
	e.From = from
	e.To = to
	e.Subject = subject
	e.HTMLBody = htmlBody
	e.MessageStream = "outbound"
	return
}

// Email represents the request body for sending an email
type Email struct {
	// Sender of email
	From string `json:"From"`
	// Receiver of email
	To string `json:"To"`
	// Additional public receivers of email
	Cc string `json:"Cc"`
	// Additional private receivers of email
	Bcc string `json:"Bcc"`

	// Email address to reply to
	ReplyTo string `json:"ReplyTo"`

	// Email subject
	Subject string `json:"Subject,omitempty"`

	// Email body as HTML
	HTMLBody string `json:"HtmlBody,omitempty"`
	// Email body as text
	TextBody string `json:"TextBody,omitempty"`

	TrackOpens bool   `json:"TrackOpens"`
	TrackLinks string `json:"TrackLinks,omitempty"`

	Attachments []Attachment `json:"Attachments,omitempty"`
	Headers     []Header     `json:"Headers,omitempty"`

	Metadata      *Metadata `json:"Metadata,omitempty"`
	MessageStream string    `json:"MessageStream,omitempty"`
	Tag           string    `json:"Tag,omitempty"`
}

// Validate will validate an email
func (e *Email) Validate() (err error) {
	var errs errors.ErrorList
	if len(e.To) == 0 {
		errs.Push(ErrEmptyToAddress)
	}

	if len(e.From) == 0 {
		errs.Push(ErrEmptyFromAddress)
	}

	if len(e.TextBody) == 0 && len(e.HTMLBody) == 0 {
		errs.Push(ErrEmptyBody)
	}

	return errs.Err()
}
