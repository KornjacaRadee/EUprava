package data

import (
	"context"
	"fmt"
	"os"
	"statistika_service/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type VehicleRegistrationRepo struct {
	cli    *mongo.Client
	logger *config.Logger
}

// Constructor which reads db configuration from environment
func New(ctx context.Context, logger *config.Logger) (*VehicleRegistrationRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &VehicleRegistrationRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (vr *VehicleRegistrationRepo) Disconnect(ctx context.Context) error {
	err := vr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (vr *VehicleRegistrationRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := vr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		vr.logger.Println(err)
	}

	// Print available databases
	databases, err := vr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		vr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (vr *VehicleRegistrationRepo) GetAll() ([]VehicleRegistration, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	registrationCollection := vr.getRegistrationCollection()

	var registrations []VehicleRegistration
	cursor, err := registrationCollection.Find(ctx, bson.M{})
	if err != nil {
		vr.logger.Println(err)
		return nil, err
	}
	if err = cursor.All(ctx, &registrations); err != nil {
		vr.logger.Println(err)
		return nil, err
	}
	return registrations, nil
}

func (vr *VehicleRegistrationRepo) GetRegistrationByID(id string) (VehicleRegistration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	registrationCollection := vr.getRegistrationCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		vr.logger.Println(err)
		return VehicleRegistration{}, err
	}

	filter := bson.M{"_id": objID}
	var registration VehicleRegistration
	err = registrationCollection.FindOne(ctx, filter).Decode(&registration)
	if err != nil {
		vr.logger.Println(err)
		return VehicleRegistration{}, err
	}

	return registration, nil
}

func (vr *VehicleRegistrationRepo) InsertRegistration(registration *VehicleRegistration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	registrationCollection := vr.getRegistrationCollection()

	result, err := registrationCollection.InsertOne(ctx, registration)
	if err != nil {
		vr.logger.Println(err)
		return err
	}
	vr.logger.Printf("Document ID: %v\n", result.InsertedID)
	return nil
}

func (vr *VehicleRegistrationRepo) UpdateRegistration(id string, registration *VehicleRegistration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	registrationCollection := vr.getRegistrationCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"registrationPlate":       registration.RegistrationPlate,
		"registrationLocation":    registration.RegistrationLocation,
		"vehicleWeight":           registration.VehicleWeight,
		"owner":                   registration.Owner,
		"fuel":                    registration.Fuel,
		"vehicleRegistrationDate": registration.VehicleRegistrationDate,
	}}
	result, err := registrationCollection.UpdateOne(ctx, filter, update)
	vr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	vr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		vr.logger.Println(err)
		return err
	}
	return nil
}

func (vr *VehicleRegistrationRepo) DeleteRegistration(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	registrationCollection := vr.getRegistrationCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := registrationCollection.DeleteOne(ctx, filter)
	if err != nil {
		vr.logger.Println(err)
		return err
	}
	vr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (vr *VehicleRegistrationRepo) getRegistrationCollection() *mongo.Collection {
	return vr.cli.Database("your_database_name").Collection("vehicle_registrations")
}
