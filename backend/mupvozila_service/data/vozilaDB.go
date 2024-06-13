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
func InsertRegisteredVehicle(vehicle *RegisterVehicle) (*mongo.InsertOneResult, error) {
	result, err := registrationCollection.InsertOne(context.Background(), vehicle)
	if err != nil {
		log.Println("Error registering vehicle:", err)
	}
	return result, err
}

// InsertVehicle registers a vehicle into the database
func InsertCar(vehicle *Car) (*mongo.InsertOneResult, error) {
	result, err := carCollection.InsertOne(context.Background(), vehicle)
	if err != nil {
		log.Println("Error inserting car into db:", err)
	}
	return result, err
}

// GetAllCars retrieves all cars from the database
func GetAllCars() ([]*Car, error) {
	var cars []*Car

	cursor, err := carCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error retrieving cars:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var car Car
		if err := cursor.Decode(&car); err != nil {
			log.Println("Error decoding car:", err)
			continue
		}
		cars = append(cars, &car)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through cars:", err)
		return nil, err
	}

	return cars, nil
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

	cursor, err := registrationCollection.Find(context.Background(), bson.M{})
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

// GetVehicleByRegistration retrieves a vehicle by its registration from the database
func GetVehicleByRegistration(registration string) (*RegisterVehicle, error) {
	if registration == "" {
		return nil, nil
	}

	var vehicle RegisterVehicle
	err := registrationCollection.FindOne(context.Background(), bson.M{"license_plate": registration}).Decode(&vehicle)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Println("Error retrieving vehicle by registration:", err)
		return nil, err
	}

	return &vehicle, nil
}

// GetAllRegistrations retrieves all registered vehicles from the database
func GetAllRegistrations() ([]*RegisterVehicle, error) {
	var vehicles []*RegisterVehicle

	cursor, err := registrationCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error retrieving registrations:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var vehicle RegisterVehicle
		if err := cursor.Decode(&vehicle); err != nil {
			log.Println("Error decoding registration:", err)
			continue
		}
		vehicles = append(vehicles, &vehicle)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through registrations:", err)
		return nil, err
	}

	return vehicles, nil
}

// GetCarsByUserID retrieves cars based on the user's ID from the database
func GetCarsByUserID(userID primitive.ObjectID) ([]*Car, error) {
	var cars []*Car

	cursor, err := carCollection.Find(context.Background(), bson.M{"owner_id": userID})
	if err != nil {
		log.Println("Error retrieving cars by user ID:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var car Car
		if err := cursor.Decode(&car); err != nil {
			log.Println("Error decoding car:", err)
			continue
		}
		cars = append(cars, &car)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through cars:", err)
		return nil, err
	}

	return cars, nil
}

// DeleteCarByID deletes a car by its ID from the database
func DeleteCarByID(carID primitive.ObjectID) error {
	_, err := carCollection.DeleteOne(context.Background(), bson.M{"_id": carID})
	if err != nil {
		log.Println("Error deleting car:", err)
		return err
	}
	return nil
}
