package postmark

import "fmt"

// Error represents an error response from the API
type Error struct {
	// ErrorCodes as specified within the Postmark API Documentation (https://postmarkapp.com/developer/api/overview#error-codes)
	ErrorCode int64 `json:"ErrorCode"`
	// Human-readable error message
	Message string `json:"Message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s <Error code: %d>", e.Message, e.ErrorCode)
}
