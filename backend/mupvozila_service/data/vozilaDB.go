package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
    "fmt"
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
func InsertRegisteredVehicle(vehicle *RegisterVehicle) (*mongo.InsertOneResult, error) {
	result, err := registrationCollection.InsertOne(context.Background(), vehicle)
	if err != nil {
		log.Println("Error registering vehicle:", err)
	}
	return result, err
}

// InsertCar registers a car into the database
func InsertCar(car *Car) (*mongo.InsertOneResult, error) {
	result, err := carCollection.InsertOne(context.Background(), car)
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

// GetCarsByOwnerJMBG retrieves cars based on the owner's JMBG from the database
func GetCarsByOwnerJMBG(ownerJMBG string) ([]*Car, error) {
	var cars []*Car

	cursor, err := carCollection.Find(context.Background(), bson.M{"owner_jmbg": ownerJMBG})
	if err != nil {
		log.Println("Error retrieving cars by owner JMBG:", err)
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


func GetCarByLicensePlate(ctx context.Context, licensePlate string) (*Car, error) {
	var car Car
	filter := bson.M{"license_plate": licensePlate}

	// Pretražujemo kolekciju vozila
	err := carCollection.FindOne(ctx, filter).Decode(&car)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("car with license plate %s not found", licensePlate)
		}
		return nil, fmt.Errorf("error finding car: %v", err)
	}

	return &car, nil
}

// GetLicensesByUserJMBG retrieves licenses by user's JMBG from the database
func GetLicensesByUserJMBG(userJMBG string) ([]*License, error) {
	var licenses []*License

	cursor, err := licenseCollection.Find(context.Background(), bson.M{"user_jmbg": userJMBG})
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

// DeleteRegistrationByID deletes a registration by its ID from the database
func DeleteRegistrationByID(registrationID primitive.ObjectID) error {
	_, err := registrationCollection.DeleteOne(context.Background(), bson.M{"_id": registrationID})
	if err != nil {
		log.Println("Error deleting registrations:", err)
		return err
	}
	return nil
}

// DeleteLicenseByID deletes a license by its ID from the database
func DeleteLicenseByID(licenseID primitive.ObjectID) error {
	_, err := licenseCollection.DeleteOne(context.Background(), bson.M{"_id": licenseID})
	if err != nil {
		log.Println("Error deleting license:", err)
		return err
	}
	return nil
}

// UpdateCar updates a car in the database
func UpdateCar(car *Car) error {
	filter := bson.M{"_id": car.ID}
	update := bson.M{
		"$set": bson.M{
			"owner_jmbg":    car.OwnerJMBG,
			"make":          car.Make,
			"model":         car.Model,
			"year":          car.Year,
			"license_plate": car.LicensePlate,
		},
	}
	_, err := carCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating car:", err)
	}
	return err
}

// UpdateRegistration updates a registration in the database
func UpdateRegistration(registration *RegisterVehicle) error {
    filter := bson.M{"_id": registration.ID}
    update := bson.M{
        "$set": bson.M{
            "car_id":           registration.CarID,
            "name":             registration.Name,
            "issuing_date":     registration.IssuingDate,
            "valid_until_date": registration.ValidUntilDate,
        },
    }
    _, err := registrationCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Println("Error updating registration:", err)
    }
    return err
}

// UpdateLicense updates a license in the database
func UpdateLicense(license *License) error {
    filter := bson.M{"_id": license.ID}
    update := bson.M{
        "$set": bson.M{
            "user_jmbg":       license.UserJMBG,
            "category":        license.Category,
            "issuing_date":    license.IssuingDate,
            "valid_until_date": license.ValidUntilDate,
            "address":         license.Address,
            "points":          license.Points,
            "is_valid":        license.IsValid,
        },
    }
    _, err := licenseCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Println("Error updating license:", err)
    }
    return err
}