package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type SaobracajnaPolicijaClient struct {
	client  *http.Client
	baseURL string
}

func NewSaobracajnaPolicijaClient(client *http.Client, baseURL string) *SaobracajnaPolicijaClient {
	return &SaobracajnaPolicijaClient{
		client:  client,
		baseURL: baseURL,
	}
}

type Prekrsaj struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`
}

type Nesreca struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Lokacija string             `bson:"lokacija,omitempty" json:"lokacija"`
	Vozilo   string             `bson:"vozilo,omitempty" json:"vozilo"`
	Vozac    string             `bson:"vozac,omitempty" json:"vozac"`
	Opis     string             `bson:"opis,omitempty" json:"opis"`
}

func (c *SaobracajnaPolicijaClient) GetAllNesreceByVozac(ctx context.Context, vozac string) ([]Nesreca, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/nesrece/vozac/%s", c.baseURL, vozac), nil)
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

	var nesrece []Nesreca
	if err := json.NewDecoder(resp.Body).Decode(&nesrece); err != nil {
		return nil, err
	}

	return nesrece, nil
}
