package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection         *mongo.Collection
	carCollection          *mongo.Collection
	licenseCollection      *mongo.Collection
	registrationCollection *mongo.Collection
)

// InitializeCollections initializes MongoDB collections
func InitializeCollections(client *mongo.Client, dbName string) {
	database := client.Database(dbName)
	userCollection = database.Collection("users")
	carCollection = database.Collection("cars")
	licenseCollection = database.Collection("licenses")
	registrationCollection = database.Collection("registrations")
}

// InsertLicense inserts a driver's license into the database
func InsertLicense(license *License) (*mongo.InsertOneResult, error) {
	result, err := licenseCollection.InsertOne(context.Background(), license)
	if err != nil {
		log.Println("Error inserting license:", err)
	}
	return result, err
}

// InsertVehicle registers a vehicle into the database
func InsertVehicle(vehicle *RegisterVehicle) (*mongo.InsertOneResult, error) {
	result, err := carCollection.InsertOne(context.Background(), vehicle)
	if err != nil {
		log.Println("Error registering vehicle:", err)
	}
	return result, err
}

// GetLicenseByID retrieves a driver's license by ID from the database
func GetLicenseByID(id primitive.ObjectID) (*License, error) {
	var license License
	err := licenseCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&license)
	if err != nil {
		log.Println("Error retrieving license:", err)
	}
	return &license, err
}

// GetVehicleByID retrieves a vehicle by ID from the database
func GetVehicleByID(id primitive.ObjectID) (*RegisterVehicle, error) {
	var vehicle RegisterVehicle
	err := carCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&vehicle)
	if err != nil {
		log.Println("Error retrieving vehicle:", err)
	}
	return &vehicle, err
}
