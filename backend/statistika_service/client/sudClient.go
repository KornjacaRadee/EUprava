package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type HouseSearchWarrant struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	IssueDate   time.Time          `bson:"issueDate" json:"issueDate"`
	DueToDate   time.Time          `bson:"dueToDate" json:"dueToDate"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
	Address     string             `bson:"address" json:"address"`
}

type Hearing struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title         string             `bson:"title" json:"title"`
	Description   string             `bson:"description" json:"description"`
	ScheduledAt   time.Time          `bson:"scheduledAt" json:"scheduledAt"`
	Duration      time.Duration      `bson:"duration" json:"duration"`
	LegalEntityID primitive.ObjectID `bson:"legalEntityId" json:"legalEntityId"`
}
type LegalRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	UserJMBG    string             `bson:"userJMBG" json:"userJMBG"`
	RequestDate time.Time          `bson:"requestDate" json:"requestDate"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
}

type LawCourtClient struct {
	baseURL string
	client  *http.Client
}

func NewLawCourtClient(baseURL string) *LawCourtClient {
	return &LawCourtClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *LawCourtClient) FetchHouseSearchWarrants(ctx context.Context) ([]HouseSearchWarrant, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/house_search_warrants/all", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch house search warrants: status %d", resp.StatusCode)
	}

	var houseSearchWarrants []HouseSearchWarrant
	if err := json.NewDecoder(resp.Body).Decode(&houseSearchWarrants); err != nil {
		return nil, err
	}

	return houseSearchWarrants, nil
}

func (c *LawCourtClient) FetchHearings(ctx context.Context) ([]Hearing, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/hearings/all", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch hearings: status %d", resp.StatusCode)
	}

	var hearings []Hearing
	if err := json.NewDecoder(resp.Body).Decode(&hearings); err != nil {
		return nil, err
	}

	return hearings, nil
}
func (c *LawCourtClient) FetchLegalRequests(ctx context.Context) ([]LegalRequest, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/legal_requests/all", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch legal requests: status %d", resp.StatusCode)
	}

	var legalRequests []LegalRequest
	if err := json.NewDecoder(resp.Body).Decode(&legalRequests); err != nil {
		return nil, err
	}

	return legalRequests, nil
}
