package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"saobracajna_policija/client"

)

type SaobracajnaPolicijaRepo struct {
	cli    *mongo.Client
	logger *log.Logger
	client *http.Client
	mupVozilaClient *client.MupvozilaClient
	authClient       *client.AuthClient
}

var (
	dbHost = os.Getenv("SAOBRACAJNA_POLICIJA_DB_HOST")
	dbPort = os.Getenv("SAOBRACAJNA_POLICIJA_DB_PORT")
	dbName = os.Getenv("SAOBRACAJNA_POLICIJA_DB_NAME")
)

func NewSaobracajnaPolicijaRepo(ctx context.Context, logger *log.Logger, mupVozilaClient *client.MupvozilaClient, authClient *client.AuthClient) *SaobracajnaPolicijaRepo {
	dburi := fmt.Sprintf("mongodb://%s:%s/", dbHost, dbPort)

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		logger.Fatal("Failed to create MongoDB client: ", err)
		return nil
	}

	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB: ", err)
		return nil
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
		},
	}

	return &SaobracajnaPolicijaRepo{
		cli:             client,
		logger:          logger,
		client:          httpClient,
		mupVozilaClient: mupVozilaClient,
		authClient:      authClient,
	}
}

func (pr *SaobracajnaPolicijaRepo) GetAllCars(ctx context.Context) ([]client.Car, error) {
    cars, err := pr.mupVozilaClient.GetAllCars()
    if err != nil {
        return nil, err
    }
    return cars, nil
}
func (pr *SaobracajnaPolicijaRepo) GetLicensesByUserJMBG(ctx context.Context, jmbg string) ([]client.License, error) {
    licenses, err := pr.mupVozilaClient.GetLicensesByUserJMBG(jmbg)
    if err != nil {
        return nil, err
    }
    return licenses, nil
}

func (r *SaobracajnaPolicijaRepo) GetUserByJMBG(ctx context.Context, jmbg string) (*client.User, error) {
    return r.authClient.GetUserByJMBG(ctx, jmbg)
}

func (pr *SaobracajnaPolicijaRepo) GetCarByLicensePlate(ctx context.Context, licensePlate string) (*client.Car, error) {
	car, err := pr.mupVozilaClient.GetCarByLicensePlate(licensePlate)
	if err != nil {
		return nil, err
	}
	return car, nil
}



func (pr *SaobracajnaPolicijaRepo) CreateNesreca(ctx context.Context, nesreca *Nesreca) error {
	collection := pr.cli.Database(dbName).Collection("nesrece")
	_, err := collection.InsertOne(ctx, nesreca)
	if err != nil {
		return err
	}
	return nil
}

func (pr *SaobracajnaPolicijaRepo) CreatePrekrsaj(ctx context.Context, prekrsaj *Prekrsaj) error {
	collection := pr.cli.Database(dbName).Collection("prekrsaji")
	_, err := collection.InsertOne(ctx, prekrsaj)
	if err != nil {
		return err
	}
	return nil
}

func (pr *SaobracajnaPolicijaRepo) GetNesrece(ctx context.Context) ([]Nesreca, error) {
	collection := pr.cli.Database(dbName).Collection("nesrece")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var nesrece []Nesreca
	if err := cursor.All(ctx, &nesrece); err != nil {
		return nil, err
	}

	return nesrece, nil
}

func (r *SaobracajnaPolicijaRepo) GetAllNesreceByVozac(ctx context.Context, vozac string) ([]Nesreca, error) {
	collection := r.cli.Database(dbName).Collection("nesrece")
	filter := bson.M{"vozac": vozac}

	var nesrece []Nesreca
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var nesreca Nesreca
		if err := cursor.Decode(&nesreca); err != nil {
			return nil, err
		}
		nesrece = append(nesrece, nesreca)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return nesrece, nil
}


func (pr *SaobracajnaPolicijaRepo) GetPrekrsaji(ctx context.Context) ([]Prekrsaj, error) {
	collection := pr.cli.Database(dbName).Collection("prekrsaji")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var prekrsaji []Prekrsaj
	if err := cursor.All(ctx, &prekrsaji); err != nil {
		return nil, err
	}

	return prekrsaji, nil
}

func (pr *SaobracajnaPolicijaRepo) DeleteNesreca(ctx context.Context, id primitive.ObjectID) error {
	collection := pr.cli.Database(dbName).Collection("nesrece")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (pr *SaobracajnaPolicijaRepo) DeletePrekrsaj(ctx context.Context, id primitive.ObjectID) error {
	collection := pr.cli.Database(dbName).Collection("prekrsaji")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
