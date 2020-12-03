package postmark

// Attachment represents an email attachment
type Attachment struct {
	// Name of attachment
	Name string `json:"Name"`
	// Content of attachment
	Content string `json:"Content"`
	// Content type of attachment
	ContentType string `json:"ContentType"`
	// Content ID of attachment
	ContentID string `json:"ContentID,omitempty"`
}
