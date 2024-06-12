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
func UpdateLicenseValidityByUserIDAndCategory(userID primitive.ObjectID, category string, isValid bool) (*mongo.UpdateResult, error) {
	filter := bson.M{"user_id": userID, "category": category}

	// Fetch current values
	cursor, err := licenseCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Error finding licenses:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	log.Println("Current IsValid values for matched licenses:")
	for cursor.Next(context.Background()) {
		var license License
		if err := cursor.Decode(&license); err != nil {
			log.Println("Error decoding license:", err)
			continue
		}
		log.Printf("License ID: %s, IsValid: %v\n", license.ID.Hex(), license.IsValid)
	}

	update := bson.M{"$set": bson.M{"is_valid": isValid}}

	result, err := licenseCollection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating license validity:", err)
		return nil, err
	}

	return result, nil
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

// GetAllLicenses retrieves all driver's licenses from the database
func GetAllLicenses() ([]*License, error) {
	var licenses []*License

	cursor, err := licenseCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error retrieving licenses:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var license License
		if err := cursor.Decode(&license); err != nil {
			log.Println("Error decoding license:", err)
			continue
		}
		licenses = append(licenses, &license)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through licenses:", err)
		return nil, err
	}

	return licenses, nil
}

// GetAllVehicles retrieves all registered vehicles from the database
func GetAllVehicles() ([]*RegisterVehicle, error) {
	var vehicles []*RegisterVehicle

	cursor, err := carCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error retrieving vehicles:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var vehicle RegisterVehicle
		if err := cursor.Decode(&vehicle); err != nil {
			log.Println("Error decoding vehicle:", err)
			continue
		}
		vehicles = append(vehicles, &vehicle)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through vehicles:", err)
		return nil, err
	}

	return vehicles, nil
}

func GetLicencesByUserID(dbClient *mongo.Client, userID primitive.ObjectID) ([]License, error) {
	licenseCollection := dbClient.Database("mupvozila_db").Collection("licences")

	var licences []License
	cursor, err := licenseCollection.Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var license License
		if err := cursor.Decode(&license); err != nil {
			return nil, err
		}
		licences = append(licences, license)
	}

	return licences, nil
}
