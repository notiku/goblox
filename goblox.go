package goblox

import (
	"net/http"

	rbxusrs "github.com/notiku/goblox/users"
)

const (
	VERSION      = "0.0.1" // Semantic versioning
	baseUsersUrl = "https://users.roblox.com/v1"
)

// Client represents a Roblox API client.
type Client struct {
	// APIKey is the key used to authenticate with the Roblox API.
	APIKey string
	// HTTPClient is the client used to make HTTP requests.
	HTTPClient *http.Client
	// BaseUsersURL is the base URL for the Roblox users API.
	BaseUsersURL string
	// Users is the users API client.
	Users *rbxusrs.Users
}

// New creates a new Roblox API client.
func New(apiKey string) *Client {
	return &Client{
		APIKey:       apiKey,
		HTTPClient:   http.DefaultClient,
		BaseUsersURL: baseUsersUrl,
		Users:        &rbxusrs.Users{APIKey: apiKey, UsersURL: baseUsersUrl},
	}
}
