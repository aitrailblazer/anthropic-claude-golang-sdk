package models

// Message represents a message in a conversation.
type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

// MessageRequest represents the request body for sending a message.
type MessageRequest struct {
	Model             string    `json:"model"`
	Messages          []Message `json:"messages"`
	MaxTokens         int       `json:"max_tokens"` // Updated to match API's 'max_tokens'
	Temperature       float64   `json:"temperature,omitempty"`
	StopSequences     []string  `json:"stop_sequences,omitempty"`
	Stream            bool      `json:"stream,omitempty"`
	AnthropicBeta     string    `json:"anthropic_beta,omitempty"`
	System            string    `json:"system,omitempty"`
	ToolChoice        interface{} `json:"tool_choice,omitempty"`
	TopK              int       `json:"top_k,omitempty"`
	TopP              float64   `json:"top_p,omitempty"`
	Metadata          interface{} `json:"metadata,omitempty"`
}

// MessageResponse represents the response from the API.
type MessageResponse struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Role       string    `json:"role"`
	Content    []Content `json:"content"`
	Model      string    `json:"model"`
	StopReason string    `json:"stop_reason"`
	StopSeq    *string   `json:"stop_sequence,omitempty"`
	Usage      Usage     `json:"usage"`
}

// Content represents a block of content in the response.
type Content struct {
	Type   string  `json:"type"`
	Text   string  `json:"text,omitempty"`
	Source *Source `json:"source,omitempty"`
}

// Source represents the source of an image.
type Source struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}

// Usage represents billing and rate-limit usage.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}
