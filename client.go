package tiktok

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// Client represents a TikTok API client
type Client struct {
	httpClient *resty.Client
	userAgent  string
	language   string
	region     string
}

// ClientOption is a function that modifies a Client
type ClientOption func(*Client)

// NewClient creates a new TikTok API client with the given options
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: resty.New().
			SetTimeout(30 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(5 * time.Second),
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		language:  "en",
		region:    "US",
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// WithUserAgent sets a custom user agent
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// WithLanguage sets the language for API requests
func WithLanguage(language string) ClientOption {
	return func(c *Client) {
		c.language = language
	}
}

// WithRegion sets the region for API requests
func WithRegion(region string) ClientOption {
	return func(c *Client) {
		c.region = region
	}
}

// Error represents a TikTok API error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
