# Anthropic Claude Golang SDK

Welcome to the Golang SDK for the Anthropic Claude API. This SDK enables developers to integrate and interact with the Anthropic Claude API to create and manage messages.

## Installation

To install the SDK, use the following command:

```bash
go get github.com/aitrailblazer/anthropic-claude-golang-sdk
```

## Usage

Here's an example of how to use the SDK to send a structured message:

```go
package main

import (
    "fmt"
    "github.com/aitrailblazer/anthropic-claude-golang-sdk"
)

func main() {
    client := anthropic.NewClient("your-api-key")
    messages := []anthropic.Message{
        {Role: "user", Content: "Hello, Claude!"},
    }
    response, err := client.SendMessage(messages, "claude-v1", 100)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Response:", response)
}
```

## Source Code Overview

### client.go

Defines the `Client` struct and `SendMessage` function:

```go
package anthropic

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
)

type Client struct {
    apiKey     string
    apiBaseURL string
}

func NewClient(apiKey string) *Client {
    return &Client{
        apiKey:     apiKey,
        apiBaseURL: "https://api.anthropic.com/v1",
    }
}

func (c *Client) SendMessage(messages []Message, model string, maxTokens int) (string, error) {
    url := c.apiBaseURL + "/messages"
    payload := map[string]interface{}{
        "model":     model,
        "messages":  messages,
        "max_tokens": maxTokens,
    }
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", fmt.Errorf("failed to marshal payload: %w", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+c.apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", fmt.Errorf("failed to decode response: %w", err)
    }

    responseContent, ok := result["content"].([]interface{})
    if !ok || len(responseContent) == 0 {
        return "", errors.New("invalid response format")
    }

    responseText, ok := responseContent[0].(map[string]interface{})["text"].(string)
    if !ok {
        return "", errors.New("invalid response text format")
    }

    return responseText, nil
}
```

### message.go

Defines the `Message` struct:

```go
package anthropic

type Message struct {
    Role    string      `json:"role"`
    Content interface{} `json:"content"`
}
```

### client_test.go

Basic test for the `SendMessage` function:

```go
package anthropic

import (
    "testing"
)

func TestSendMessage(t *testing.T) {
    client := NewClient("test-api-key")

    _, err := client.SendMessage([]Message{{Role: "user", Content: "Test message"}}, "claude-v1", 100)
    if err != nil {
        t.Errorf("SendMessage failed: %v", err)
    }
}
```

## Feedback

We welcome feedback to improve the SDK and documentation. Please submit any issues or suggestions via the GitHub repository.

---

Thank you for using the Anthropic Claude Golang SDK! We hope it enhances your development experience.