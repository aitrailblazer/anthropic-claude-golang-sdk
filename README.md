# Anthropic Claude Golang SDK

A Go SDK for interacting with the [Anthropic Claude API](https://docs.anthropic.com/en/api/messages). This SDK allows you to send structured messages and handle responses from Anthropic's conversational AI models seamlessly within your Go applications.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Using the Bash Script](#using-the-bash-script)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Running the Example](#running-the-example)
- [Project Structure](#project-structure)
- [API Reference](#api-reference)
- [Logging](#logging)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Easy Integration**: Simplifies communication with Anthropic's Claude API using Go.
- **Environment Variable Configuration**: Securely manage your API keys and settings.
- **Detailed Logging**: Helps in debugging by logging request and response details.
- **Flexible Configuration**: Easily adjust parameters like `max_tokens`, `temperature`, and model selection via environment variables.

## Prerequisites

- **Go**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
- **Anthropic API Key**: Obtain your API key from the [Anthropic Console](https://console.anthropic.com/).

## Installation

### Using the Bash Script

A bash script is provided to automate the setup of the SDK, including creating the project structure, initializing the Go module, and generating essential files.

1. **Save the Script**

   Copy the script below and save it as `create_anthropic_sdk.sh` in your desired directory.

   ```bash
   #!/bin/bash
   
   # Exit immediately if a command exits with a non-zero status
   set -e
   
   # Set the project directory name
   PROJECT_DIR="anthropic-claude-golang-sdk"
   
   # Check if the project directory already exists
   if [ -d "$PROJECT_DIR" ]; then
       echo "Directory '$PROJECT_DIR' already exists. Please remove it or choose a different project name."
       exit 1
   fi
   
   # Create the main project directory and subdirectories
   echo "Creating project directory structure..."
   mkdir -p "$PROJECT_DIR/api" "$PROJECT_DIR/models" "$PROJECT_DIR/examples"
   
   # Navigate to the project directory
   cd "$PROJECT_DIR" || exit
   
   # Initialize Go module
   echo "Initializing Go module..."
   go mod init github.com/aitrailblazer/anthropic-claude-golang-sdk
   
   # Create client.go in the api directory with logging and correct headers
   echo "Creating api/client.go..."
   cat > api/client.go <<EOL
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
   EOL
   
   # Create message.go in the api directory with SendMessage function
   echo "Creating api/message.go..."
   cat > api/message.go <<EOL
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
   EOL
   
   # Create message.go in the models directory with request and response structures
   echo "Creating models/message.go..."
   cat > models/message.go <<EOL
   package models
   
   // Message represents a message in a conversation.
   type Message struct {
   	Role    string      \`json:"role"\`
   	Content interface{} \`json:"content"\`
   }
   
   // MessageRequest represents the request body for sending a message.
   type MessageRequest struct {
   	Model             string     \`json:"model"\`
   	Messages          []Message  \`json:"messages"\`
   	MaxTokens         int        \`json:"max_tokens"\` // Updated to match API's 'max_tokens'
   	Temperature       float64    \`json:"temperature,omitempty"\`
   	StopSequences     []string   \`json:"stop_sequences,omitempty"\`
   	Stream            bool       \`json:"stream,omitempty"\`
   	AnthropicBeta     string     \`json:"anthropic_beta,omitempty"\`
   	System            string     \`json:"system,omitempty"\`
   	ToolChoice        interface{} \`json:"tool_choice,omitempty"\`
   	TopK              int        \`json:"top_k,omitempty"\`
   	TopP              float64    \`json:"top_p,omitempty"\`
   	Metadata          interface{} \`json:"metadata,omitempty"\`
   }
   
   // MessageResponse represents the response from the API.
   type MessageResponse struct {
   	ID         string    \`json:"id"\`
   	Type       string    \`json:"type"\`
   	Role       string    \`json:"role"\`
   	Content    []Content \`json:"content"\`
   	Model      string    \`json:"model"\`
   	StopReason string    \`json:"stop_reason"\`
   	StopSeq    *string   \`json:"stop_sequence,omitempty"\`
   	Usage      Usage     \`json:"usage"\`
   }
   
   // Content represents a block of content in the response.
   type Content struct {
   	Type   string  \`json:"type"\`
   	Text   string  \`json:"text,omitempty"\`
   	Source *Source \`json:"source,omitempty"\`
   }
   
   // Source represents the source of an image.
   type Source struct {
   	Type      string \`json:"type"\`
   	MediaType string \`json:"media_type"\`
   	Data      string \`json:"data"\`
   }
   
   // Usage represents billing and rate-limit usage.
   type Usage struct {
   	InputTokens  int \`json:"input_tokens"\`
   	OutputTokens int \`json:"output_tokens"\`
   }
   EOL
   
   # Create example.go in the examples directory with detailed logging and correct fields
   echo "Creating examples/example.go..."
   cat > examples/example.go <<EOL
   package main
   
   import (
   	"fmt"
   	"log"
   	"os"
   	"strconv"
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
   
   	// Get the model name from environment variable or use default
   	model := os.Getenv("ANTHROPIC_MODEL")
   	if model == "" {
   		model = "claude-1.3" // Default model
   	}
   
   	// Get MaxTokens from environment variable or use default
   	maxTokensStr := os.Getenv("ANTHROPIC_MAX_TOKENS")
   	maxTokens := 300 // Default value
   	if maxTokensStr != "" {
   		if val, err := strconv.Atoi(maxTokensStr); err == nil {
   			maxTokens = val
   		} else {
   			log.Printf("Invalid ANTHROPIC_MAX_TOKENS value '%s'. Using default %d.", maxTokensStr, maxTokens)
   		}
   	}
   
   	// Prepare the message request with conversational turns
   	request := models.MessageRequest{
   		Model: model, // Use the model name from environment variable
   		Messages: []models.Message{
   			{Role: "user", Content: "Hello there."},
   			{Role: "assistant", Content: "Hi, I'm Claude. How can I help you?"},
   			{Role: "user", Content: "Can you explain LLMs in plain English?"},
   		},
   		MaxTokens:   maxTokens, // Set from environment variable
   		Temperature: 0.7,        // Creativity control
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
   EOL
   
   # Run go mod tidy to resolve dependencies
   echo "Running 'go mod tidy' to resolve dependencies..."
   go mod tidy
   
   echo "========================================"
   echo "Project setup complete!"
   echo "========================================"
   echo "Next Steps:"
   echo "1. Set your ANTHROPIC_API_KEY environment variable."
   echo "   - On Linux/macOS:"
   echo "       export ANTHROPIC_API_KEY=\"your-api-key\""
   echo "   - On Windows (PowerShell):"
   echo "       \$env:ANTHROPIC_API_KEY=\"your-api-key\""
   echo ""
   echo "2. Optionally, set the ANTHROPIC_MODEL and ANTHROPIC_MAX_TOKENS environment variables."
   echo "   - On Linux/macOS:"
   echo "       export ANTHROPIC_MODEL=\"claude-1.3\""
   echo "       export ANTHROPIC_MAX_TOKENS=300"
   echo "   - On Windows (PowerShell):"
   echo "       \$env:ANTHROPIC_MODEL=\"claude-1.3\""
   echo "       \$env:ANTHROPIC_MAX_TOKENS=300"
   echo ""
   echo "3. Navigate to the examples directory:"
   echo "       cd examples"
   echo ""
   echo "4. Run the example:"
   echo "       go run example.go"
   echo "========================================"
