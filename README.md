# Anthropic Claude Golang SDK

A Go SDK for interacting with the [Anthropic Claude API](https://console.anthropic.com/). This SDK allows you to send structured messages and handle responses from Anthropic's conversational AI models seamlessly within your Go applications.

## Table of Contents

- [Anthropic Claude Golang SDK](#anthropic-claude-golang-sdk)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Cloning the Repository](#cloning-the-repository)
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

### Cloning the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/aitrailblazer/anthropic-claude-golang-sdk.git
## Configuration

The SDK relies on environment variables for configuration to ensure security and flexibility.

- `ANTHROPIC_API_KEY` (required): Your unique API key for authenticating with Anthropic's API.
- `ANTHROPIC_MODEL` (optional): The model name to use (default is `claude-1.3`).
- `ANTHROPIC_MAX_TOKENS` (optional): Maximum number of tokens to generate in the response (default is 300).

### Setting Environment Variables

On Linux/macOS:

```bash
export ANTHROPIC_API_KEY="your-api-key"
export ANTHROPIC_MODEL="claude-1.3"         # Optional
export ANTHROPIC_MAX_TOKENS=300             # Optional
```

## Usage

### Running the Example

An example program is provided to demonstrate how to use the SDK to send messages and handle responses.

Navigate to the Examples Directory:

```bash
cd anthropic-claude-golang-sdk/examples
```

Run the Example:

```bash
go run example.go
```

Example Output:

```plaintext
2024/09/18 00:28:55 Sending POST request to https://api.anthropic.com/v1/messages with body: {"model":"claude-1.3","messages":[{"role":"user","content":"Hello there."},{"role":"assistant","content":"Hi, I'm Claude. How can I help you?"},{"role":"user","content":"Can you explain LLMs in plain English?"}],"max_tokens":100,"temperature":0.7}
2024/09/18 00:28:57 Received response with status: 200 OK
2024/09/18 00:28:57 Response body: {"id":"msg_01YNjiU5mFbKVy8MNHCP7brd","type":"message","role":"assistant","model":"claude-1.3","content":[{"type":"text","text":"Sure, here's a simple explanation of large language models or LLMs:\n\n• LLMs are AI systems that have been trained on massive amounts of text data. They learn patterns and relationships between words and phrases to understand language.\n\n• The more text data they are trained on, the more they can understand the nuances, complexities and ambiguities of natural language. Some recent LLMs have been trained on hundreds of billions of words.\n\n• LLMs can be use"}],"stop_reason":"max_tokens","stop_sequence":null,"usage":{"input_tokens":41,"output_tokens":100}}
Response: Sure, here's a simple explanation of large language models or LLMs:

• LLMs are AI systems that have been trained on massive amounts of text data. They learn patterns and relationships between words and phrases to understand language.

• The more text data they are trained on, the more they can understand the nuances, complexities and ambiguities of natural language. Some recent LLMs have been trained on hundreds of billions of words.
```

## Project Structure

```plaintext
anthropic-claude-golang-sdk/
│
├── api
│   ├── client.go     # Handles API requests, authentication, and logging
│   └── message.go    # Contains the SendMessage function to interact with the Messages API
│
├── models
│   └── message.go    # Defines the request and response structures
├── examples
│   └── example.go    # Demonstrates how to use the SDK
├── go.mod            # Go module file
└── go.sum            # Go module checksum file
```

## API Reference

Refer to the official [API documentation](https://console.anthropic.com/) for detailed information on available endpoints and parameters.

## Logging

The SDK provides detailed logging to assist with debugging. Logs include request and response details.

## Error Handling

Comprehensive error handling is implemented to manage API errors and connection issues effectively.

## Contributing

Contributions are welcome! Please refer to the [contributing guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
