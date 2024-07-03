package client

import (

    "encoding/json"
    "fmt"
    "net/http"
)

type MupvozilaClient struct {
    client  *http.Client
    baseURL string
}

func NewMupvozilaClient(client *http.Client, baseURL string) *MupvozilaClient {
	return &MupvozilaClient{
		client:  client,
		baseURL: baseURL,
	}
}

type License struct {
    ID             string `json:"id"`
    UserJMBG       string `json:"user_jmbg"`
    Category       string `json:"category"`
    IssuingDate    string `json:"issuing_date"`
    ValidUntilDate string `json:"valid_until_date"`
    Address        string `json:"address"`
    Points         int    `json:"points"`
    IsValid        bool   `json:"is_valid"`
}

type Car struct {
    ID           string `json:"id"`
    OwnerJMBG    string `json:"owner_jmbg"`
    Make         string `json:"make"`
    Model        string `json:"model"`
    Year         int    `json:"year"`
    LicensePlate string `json:"license_plate"`
}

// GetLicensesByUserJMBG retrieves licenses for a user by their JMBG.
func (c *MupvozilaClient) GetLicensesByUserJMBG(jmbg string) ([]License, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/licenses/user/%s", c.baseURL, jmbg), nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
    }

    var licenses []License
    if err := json.NewDecoder(resp.Body).Decode(&licenses); err != nil {
        return nil, err
    }

    return licenses, nil
}

// GetAllCars retrieves all cars from the MupVozila service.
func (c *MupvozilaClient) GetAllCars() ([]Car, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/getAllCars", c.baseURL), nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
    }

    var cars []Car
    if err := json.NewDecoder(resp.Body).Decode(&cars); err != nil {
        return nil, err
    }

    return cars, nil
}
// GetCarByLicensePlate retrieves a car by its license plate.
func (c *MupvozilaClient) GetCarByLicensePlate(licensePlate string) (*Car, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/car/carPlate/%s", c.baseURL, licensePlate), nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
    }

    var car Car
    if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
        return nil, err
    }

    return &car, nil
}
