package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// MupvozilaClient represents a client for the mupvozila service.
type MupvozilaClient struct {
	client  *http.Client
	baseURL string
}

// NewMupvozilaClient creates a new instance of MupvozilaClient.
func NewMupvozilaClient(client *http.Client, baseURL string) *MupvozilaClient {
	return &MupvozilaClient{
		client:  client,
		baseURL: baseURL,
	}
}

// UpdateLicenceValidity updates the licence validity for a legal entity.
func (c *MupvozilaClient) UpdateLicenceValidity(ctx context.Context, userId string, category string) error {
	payload := map[string]bool{
		"is_valid": false,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/updateLicenseValidity/user/%s/category/%s", c.baseURL, userId, category), bytes.NewReader(payloadBytes))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
