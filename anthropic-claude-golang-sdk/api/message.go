package api

import (
	"encoding/json"
	"fmt"

	"github.com/aitrailblazer/anthropic-claude-golang-sdk/models"
)

// SendMessage sends a message to the Anthropic API and returns the response.
func (c *Client) SendMessage(request models.MessageRequest) (*models.MessageResponse, error) {
	respBody, err := c.MakeRequest("POST", "messages", request)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}

	var response models.MessageResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &response, nil
}
