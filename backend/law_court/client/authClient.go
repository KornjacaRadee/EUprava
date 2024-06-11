package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthClient represents a client for the auth service.
type AuthClient struct {
	client  *http.Client
	baseURL string
}

// NewAuthClient creates a new instance of AuthClient.
func NewAuthClient(client *http.Client, baseURL string) *AuthClient {
	return &AuthClient{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *AuthClient) GetUserByID(ctx context.Context, userID string) (*User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/users/%s", c.baseURL, userID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *AuthClient) GetUserByJMBG(ctx context.Context, jmbg string) (*User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/users/jmbg/%s", c.baseURL, jmbg), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
