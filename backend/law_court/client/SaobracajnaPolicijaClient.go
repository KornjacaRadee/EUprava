package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Violation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`
}

type SaobracajnaPolicijaClient struct {
	baseURL string
	client  *http.Client
}

func NewSaobracajnaPolicijaClient(baseURL string) *SaobracajnaPolicijaClient {
	return &SaobracajnaPolicijaClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *SaobracajnaPolicijaClient) FetchPrekrsaji(ctx context.Context) ([]Violation, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/nesreca", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch prekrsaji: status %d", resp.StatusCode)
	}

	var prekrsaji []Violation
	if err := json.NewDecoder(resp.Body).Decode(&prekrsaji); err != nil {
		return nil, err
	}

	return prekrsaji, nil
}
