package postmark

// Header represents a header value
type Header struct {
	// Name of header
	Name string `json:"Name"`
	// Header value
	Value string `json:"Value"`
}
