package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Client struct to hold the API key and other client-specific information.
type Client struct {
	APIKey string
}

// NewClient creates a new client by reading the API key from the environment variable.
func NewClient() (*Client, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API key not found. Please set the ANTHROPIC_API_KEY environment variable")
	}
	return &Client{APIKey: apiKey}, nil
}

// MakeRequest sends an HTTP request to the Anthropic API and returns the response.
func (c *Client) MakeRequest(method, endpoint string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("https://api.anthropic.com/v1/%s", endpoint)

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", err)
	}

	// Log the request details
	log.Printf("Sending %s request to %s with body: %s\n", method, url, string(jsonData))

	// Create the HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01") // Update this if a different version is needed

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Log the response status and body
	log.Printf("Received response with status: %s\n", resp.Status)
	log.Printf("Response body: %s\n", string(respBody))

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from server: %s, %s", resp.Status, string(respBody))
	}

	return respBody, nil
}
