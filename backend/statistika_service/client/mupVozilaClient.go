package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// Car represents a car entity
type Car struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerJMBG    string             `bson:"owner_jmbg" json:"owner_jmbg" validate:"required"`
	Make         string             `bson:"make" json:"make" validate:"required"`
	Model        string             `bson:"model" json:"model" validate:"required"`
	Year         int                `bson:"year" json:"year" validate:"required"`
	LicensePlate string             `bson:"license_plate" json:"license_plate" validate:"required"`
}

type Cars []*Car

// License represents a driver's license
type License struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserJMBG       string             `bson:"user_jmbg" json:"user_jmbg" validate:"required"`
	Category       string             `bson:"category" json:"category" validate:"required"`
	IssuingDate    time.Time          `bson:"issuing_date" json:"issuing_date" validate:"required"`
	ValidUntilDate time.Time          `bson:"valid_until_date" json:"valid_until_date" validate:"required"`
	Address        string             `bson:"address" json:"address"`
	Points         int                `bson:"points" json:"points"`
	IsValid        bool               `bson:"is_valid" json:"is_valid"`
}

type Licenses []*License

// RegisterVehicle represents a registered vehicle
type RegisterVehicle struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CarID          primitive.ObjectID `bson:"car_id" json:"car_id" validate:"required"`
	Name           string             `bson:"name" json:"name" validate:"required"`
	IssuingDate    time.Time          `bson:"issuing_date" json:"issuing_date" validate:"required"`
	ValidUntilDate time.Time          `bson:"valid_until_date" json:"valid_until_date" validate:"required"`
}

// MupVozilaClient is a client for interacting with the MUP Vozila service
type MupVozilaClient struct {
	baseURL string
	client  *http.Client
}

// NewMupVozilaClient creates a new MupVozilaClient
func NewMupVozilaClient(baseURL string) *MupVozilaClient {
	return &MupVozilaClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// FetchCars fetches the list of cars
func (c *MupVozilaClient) FetchCars(ctx context.Context) ([]Car, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/getAllCars", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch cars: status %d", resp.StatusCode)
	}

	var cars []Car
	if err := json.NewDecoder(resp.Body).Decode(&cars); err != nil {
		return nil, err
	}

	return cars, nil
}

// FetchLicenses fetches the list of licenses
func (c *MupVozilaClient) FetchLicenses(ctx context.Context) ([]License, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/getAllLicenses", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch licenses: status %d", resp.StatusCode)
	}

	var licenses []License
	if err := json.NewDecoder(resp.Body).Decode(&licenses); err != nil {
		return nil, err
	}

	return licenses, nil
}

// FetchRegisteredVehicles fetches the list of registered vehicles
func (c *MupVozilaClient) FetchRegisteredVehicles(ctx context.Context) ([]RegisterVehicle, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/registrations", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch registered vehicles: status %d", resp.StatusCode)
	}

	var registeredVehicles []RegisterVehicle
	if err := json.NewDecoder(resp.Body).Decode(&registeredVehicles); err != nil {
		return nil, err
	}

	return registeredVehicles, nil
}
