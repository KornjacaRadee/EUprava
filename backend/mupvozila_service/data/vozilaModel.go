package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Car represents a car in the system
type Car struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerID      primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	Make         string             `bson:"make" json:"make" validate:"required"`
	Model        string             `bson:"model" json:"model" validate:"required"`
	Year         int                `bson:"year" json:"year" validate:"required"`
	LicensePlate string             `bson:"license_plate" json:"license_plate" validate:"required"`
}

type Cars []*Car

// License represents a driver's license
type License struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"user_id"`
	Category       string             `bson:"category" json:"category" validate:"required"`
	IssuingDate    time.Time          `bson:"issuing_date" json:"issuing_date" validate:"required"`
	ValidUntilDate time.Time          `bson:"valid_until_date" json:"valid_until_date" validate:"required"`
	Address        string             `bson:"address" json:"address"`
	Points         int                `bson:"points" json:"points"`
	IsValid        bool               `bson:"is_valid" json:"is_valid"`
}

type Licences []*License

// RegisterVehicle represents a registered vehicle
type RegisterVehicle struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CarID          primitive.ObjectID `bson:"car_id" json:"car_id" validate:"required"`
	Name           string             `bson:"name" json:"name" validate:"required"`
	IssuingDate    time.Time          `bson:"issuing_date" json:"issuing_date" validate:"required"`
	ValidUntilDate time.Time          `bson:"valid_until_date" json:"valid_until_date" validate:"required"`
}
