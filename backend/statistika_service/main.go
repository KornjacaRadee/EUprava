package main

import (
	"context"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"statistika_service/config"
	"statistika_service/data"
	handlers "statistika_service/vehicleRegistrationHandlers"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8083"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	//logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	//storeLogger := log.New(os.Stdout, "[accommodation-store] ", log.LstdFlags)
	//imageStorageLogger := log.New(os.Stdout, "[accommodation-image_storage] ", log.LstdFlags)
	//redisLogger := log.New(os.Stdout, "[accommodation-cache] ", log.LstdFlags)
	logger := config.NewLogger("./logging/log.log")
	// NoSQL: Initialize Accommodation Repository store
	store, err := data.New(timeoutContext, logger)
	if err != nil {
		logger.Fatalf("Failed initializing Accommodation Repository Store", err)
	}
	defer store.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	store.Ping()

	//Initialize the handler and inject logger
	vehicleRegistrationHandler := handlers.NewVehicleRegistrationHandler(logger, store)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(vehicleRegistrationHandler.MiddlewareContentTypeSet)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/all", vehicleRegistrationHandler.GetAllVehicles)
	//getRouter.HandleFunc("/accommodation/walk", vehicleRegistrationHandler.WalkRoot)
	getRouter.HandleFunc("/{id}", vehicleRegistrationHandler.GetVehicle)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/new", vehicleRegistrationHandler.PostVehicleRegistration)
	postRouter.Use(vehicleRegistrationHandler.MiddlewareVehicleDeserialization)

	patchRouter := router.Methods(http.MethodPatch).Subrouter()
	patchRouter.HandleFunc("/patch/{id}", vehicleRegistrationHandler.PatchVehicleRegistration)
	patchRouter.Use(vehicleRegistrationHandler.MiddlewareVehicleDeserialization)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/delete/{id}", vehicleRegistrationHandler.DeleteVehicleRegistration)
	// ...

	// Start server

	//router.Use(handlers2.AuthMiddleware)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatalf("Cannot distribute all the connections to goroutines", err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shut down gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatalf("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
