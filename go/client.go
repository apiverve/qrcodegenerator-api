// Package qrcodegenerator provides a Go client for the QR Code Generator API.
//
// For more information, visit: https://apiverve.com/marketplace/qrcodegenerator?utm_source=go&utm_medium=readme
package qrcodegenerator

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL        = "https://api.apiverve.com/v1/qrcodegenerator"
	defaultTimeout = 30 * time.Second
)

// Client is the QR Code Generator API client.
type Client struct {
	apiKey     string
	httpClient *resty.Client
}

// NewClient creates a new QR Code Generator API client.
func NewClient(apiKey string) *Client {
	client := resty.New()
	client.SetTimeout(defaultTimeout)
	client.SetHeader("Content-Type", "application/json")

	return &Client{
		apiKey:     apiKey,
		httpClient: client,
	}
}

// SetTimeout sets the HTTP client timeout.
func (c *Client) SetTimeout(timeout time.Duration) {
	c.httpClient.SetTimeout(timeout)
}


// Execute makes a request to the QR Code Generator API with typed parameters.
//
// Parameters are validated before sending the request. If validation fails,
// an error is returned immediately without making a network request.
//
// Available parameters:
//   - value (required): string - The text or data to encode in the QR code
//   - type: string - The type of data being encoded. Advanced types (wifi, vcard) are premium.
//   - format: string - Output format. Vector formats (svg, webp) are premium.
//   - size: number - Size of the QR code in pixels (50-2048)
//   - margin: number - Margin around the QR code in pixels (0-100)
//   - color: string - Foreground color as hex code (e.g., #000000)
//   - backgroundColor: string - Background color as hex code (e.g., #ffffff)
//   - dotStyle: string - Style of QR code dots
//   - cornerSquareStyle: string - Style of corner squares
//   - cornerDotStyle: string - Style of corner dots
//   - gradient: object - Gradient configuration with type (linear, radial) and colorStops array
//   - logo: string - URL of logo image to place in center of QR code [format: url]
//   - logoSize: number - Size of logo relative to QR code (0.1-0.5)
//   - logoMargin: number - Margin around logo in pixels
func (c *Client) Execute(req *Request) (*Response, error) {
	if c.apiKey == "" {
		return nil, errors.New("API key is required. Get your API key at: https://apiverve.com")
	}

	// Validate parameters before making request
	if req != nil {
		if err := req.Validate(); err != nil {
			return nil, err
		}
	}

	request := c.httpClient.R().
		SetHeader("x-api-key", c.apiKey).
		SetResult(&Response{}).
		SetError(&ErrorResponse{})

	// POST request with JSON body
	resp, err := request.
		SetBody(req).
		Post(baseURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		if errResp, ok := resp.Error().(*ErrorResponse); ok {
			return nil, fmt.Errorf("API error: %s", errResp.Error)
		}
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*Response)
	if !ok {
		return nil, errors.New("failed to parse response")
	}

	return result, nil
}

// ExecuteRaw makes a request with a raw map of parameters (for dynamic usage).
func (c *Client) ExecuteRaw(params map[string]interface{}) (*Response, error) {
	if c.apiKey == "" {
		return nil, errors.New("API key is required. Get your API key at: https://apiverve.com")
	}

	request := c.httpClient.R().
		SetHeader("x-api-key", c.apiKey).
		SetResult(&Response{}).
		SetError(&ErrorResponse{})

	resp, err := request.
		SetBody(params).
		Post(baseURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		if errResp, ok := resp.Error().(*ErrorResponse); ok {
			return nil, fmt.Errorf("API error: %s", errResp.Error)
		}
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode())
	}

	result, ok := resp.Result().(*Response)
	if !ok {
		return nil, errors.New("failed to parse response")
	}

	return result, nil
}

// Response represents the API response.
type Response struct {
	Status string       `json:"status"`
	Error  interface{}  `json:"error"`
	Data   ResponseData `json:"data"`
}

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
