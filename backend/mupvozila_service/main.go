package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mupvozila_service/config"
	"mupvozila_service/data"
	mupVozilaHandlers "mupvozila_service/mupVozilaHandlers" // Import handlers package
)

func main() {
	// MongoDB connection setup
	dbURI := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(dbURI)
	dbClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Disconnect(context.TODO())

	// Initialize MongoDB collections
	data.InitializeCollections(dbClient, "mupVozilaDB")

	// Logger setup
	logger := config.NewLogger("./logging/log.log")

	// Server setup
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/licenses", mupVozilaHandlers.IssueLicenseHandler).Methods("POST")
	r.HandleFunc("/vehicles", mupVozilaHandlers.RegisterVehicleHandler).Methods("POST")
	r.HandleFunc("/cars", mupVozilaHandlers.InsertCarHandler).Methods("POST")
	r.HandleFunc("/cars/user/{user_id}", mupVozilaHandlers.GetCarsByUserIDHandler).Methods("GET")
	r.HandleFunc("/cars/{car_id}", mupVozilaHandlers.DeleteCarByIDHandler).Methods("DELETE")
	r.HandleFunc("/cars/{id}", mupVozilaHandlers.UpdateCarHandler).Methods("PUT")
	r.HandleFunc("/getAllLicenses", mupVozilaHandlers.GetAllLicensesHandler).Methods("GET")
	r.HandleFunc("/licenses/{license_id}", mupVozilaHandlers.DeleteLicenseByIDHandler).Methods("DELETE")
	r.HandleFunc("/licenses/user/{jmbg}", mupVozilaHandlers.GetLicensesByUserJMBGHandler).Methods("GET")
	r.HandleFunc("/getAllVehicles", mupVozilaHandlers.GetAllVehiclesHandler).Methods("GET")
	r.HandleFunc("/getAllCars", mupVozilaHandlers.GetAllCarsHandler).Methods("GET")
	r.HandleFunc("/getLicenseById/{id}", mupVozilaHandlers.GetLicenseByIDHandler).Methods("GET")
	r.HandleFunc("/getVehicleById/{id}", mupVozilaHandlers.GetVehicleByIDHandler).Methods("GET")
	r.HandleFunc("/getLicencesByUserID/user/{id}", mupVozilaHandlers.GetLicencesByUserID(dbClient)).Methods("GET")
	// New route for retrieving vehicle by registration
	r.HandleFunc("/vehicles/registration/{registration}", mupVozilaHandlers.GetVehicleByRegistrationHandler).Methods("GET")
	// New route for retrieving all registrations
	r.HandleFunc("/registrations", mupVozilaHandlers.GetAllRegistrationsHandler).Methods("GET")
	r.HandleFunc("/registrations/{id}", mupVozilaHandlers.DeleteRegistrationHandler).Methods("DELETE")
	r.HandleFunc("/vehicles/register", mupVozilaHandlers.RegisterVehicleHandler).Methods("POST")

 // New routes for updating registrations and licenses
    r.HandleFunc("/registrations/{id}", mupVozilaHandlers.UpdateRegistrationHandler).Methods("PUT")
    r.HandleFunc("/licenses/{id}", mupVozilaHandlers.UpdateLicenseHandler).Methods("PUT")


	// CORS setup
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	handlerWithCORS := handlers.CORS(headers, methods, origins)(r)

	// Start server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handlerWithCORS,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		logger.Println("Server listening on port", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, os.Kill)
	<-stop

	logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Error shutting down server: %s\n", err)
	}

	logger.Println("Server stopped gracefully")
}
