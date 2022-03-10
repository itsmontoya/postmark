package postmark

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	host = "https://api.postmarkapp.com"

	endpointEmail             = "/email"
	endpointEmailWithTemplate = "/email/withTemplate"
)

var hostURL = parseURL(host)

// New will initialize and return a new Client
func New(apiKey string) *Client {
	// Create new client
	c := makeClient(apiKey)
	// Return pointer to created client
	return &c
}

func makeClient(apiKey string) (c Client) {
	// Set API key
	c.apiKey = apiKey
	return
}

// Client manages the request to the Postmark API
type Client struct {
	hc http.Client
	// Server token
	apiKey string
}

func (c *Client) request(method, endpoint string, body io.Reader, respValue interface{}) (err error) {
	var req *http.Request
	// Get URL for target endpoint
	url := getURL(endpoint)
	// Initialize a new HTTP request value
	if req, err = http.NewRequest(method, url, body); err != nil {
		return
	}

	// Set server token
	req.Header.Set("X-Postmark-Server-Token", c.apiKey)
	// Set content type
	req.Header.Set("Content-Type", "application/json")
	// Set accept
	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	// Make request
	if resp, err = c.hc.Do(req); err != nil {
		return
	}
	// Close the response body after this function has exited
	defer resp.Body.Close()

	// Check status code to ensure response is OK
	if resp.StatusCode >= 400 {
		// We've received an error status code, parse the response body as an error and return
		err = handleError(resp.Body)
		return
	}

	// Decode response body as provided response value
	return json.NewDecoder(resp.Body).Decode(respValue)
}

// Email will send an email
func (c *Client) Email(e Email) (resp EmailResponse, err error) {
	// Initialize buffer to write request body to
	buf := bytes.NewBuffer(nil)
	// Write email request as JSON to the buffer
	if err = json.NewEncoder(buf).Encode(e); err != nil {
		return
	}

	// Make POST request to the email endpoint with the request buffer as the body
	err = c.request("POST", endpointEmail, buf, &resp)
	return
}

// Email will send an email
func (c *Client) EmailWithTemplate(e EmailWithTemplate) (resp EmailResponse, err error) {
	// Initialize buffer to write request body to
	buf := bytes.NewBuffer(nil)
	// Write email request as JSON to the buffer
	if err = json.NewEncoder(buf).Encode(e); err != nil {
		return
	}

	// Make POST request to the email endpoint with the request buffer as the body
	err = c.request("POST", endpointEmailWithTemplate, buf, &resp)
	return
}
