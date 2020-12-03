package postmark

import (
	"encoding/json"
	"io"
	"net/url"
)

func parseURL(urlStr string) url.URL {
	// Parse string as URL
	parsed, _ := url.Parse(urlStr)
	// Return a de-referenced url.URL
	return *parsed
}

func getURL(endpoint string) (url string) {
	// Copy host URL
	u := hostURL
	// Set path as the provided endpoint
	u.Path = endpoint
	// Return string version of url.RUL
	return u.String()
}

func handleError(r io.Reader) (err error) {
	var e Error
	// Decode reader as an Error
	if err = json.NewDecoder(r).Decode(&e); err != nil {
		// Error encountered while decoding reader as Error
		return
	}

	// Return error
	return e
}
