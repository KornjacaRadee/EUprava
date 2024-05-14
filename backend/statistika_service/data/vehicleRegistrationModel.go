package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type VehicleRegistration struct {
	Id                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RegistrationPlate       string             `json:"registrationPlate" bson:"registrationPlate"`
	RegistrationLocation    string             `json:"registrationLocation" bson:"registrationLocation"`
	VehicleWeight           float64            `json:"vehicleWeight" bson:"vehicleWeight"`
	Owner                   string             `json:"owner" bson:"owner"`
	Fuel                    string             `json:"fuel" bson:"fuel"`
	VehicleRegistrationDate time.Time          `json:"vehicleRegistrationDate" bson:"vehicleRegistrationDate"`
}
type VehicleRegistrarion []*VehicleRegistration

// Functions to encode and decode products to json and from json.
func (v *VehicleRegistration) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(v)
}
func (a VehicleRegistrarion) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (v *VehicleRegistration) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(v)
}
