package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aitrailblazer/anthropic-claude-golang-sdk/api"
	"github.com/aitrailblazer/anthropic-claude-golang-sdk/models"
)

// extractText extracts and concatenates text from the Content blocks.
func extractText(contents []models.Content) string {
	var texts []string
	for _, content := range contents {
		if content.Type == "text" {
			texts = append(texts, content.Text)
		}
	}
	return strings.Join(texts, " ")
}

func main() {
	// Initialize the SDK client (API key is read from environment variable)
	client, err := api.NewClient()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Prepare the message request with conversational turns
	request := models.MessageRequest{
		Model: "claude-1.3", // Updated to the correct model name
		Messages: []models.Message{
			{Role: "user", Content: "Hello there."},
			{Role: "assistant", Content: "Hi, I'm Claude. How can I help you?"},
			{Role: "user", Content: "Can you explain LLMs in plain English?"},
		},
		MaxTokens:   100, // This is required to specify the max number of tokens
		Temperature: 0.7, // Creativity control
	}

	// Send the message request and get the response
	response, err := client.SendMessage(request)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// Extract and print the response content
	responseText := extractText(response.Content)
	fmt.Println("Response:", responseText)
}
