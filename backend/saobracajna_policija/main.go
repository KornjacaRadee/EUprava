package main

import (
    "context"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "net/http"
    "os"
    "os/signal"
    "saobracajna_policija/data"
    "saobracajna_policija/handlers" // Ensure this import path is correct
    "saobracajna_policija/client"
    "time"
)

func main() {
    // Initialize logger
    logger := log.New(os.Stdout, "[saobracajna-policija] ", log.LstdFlags)

    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8084"
    }

    // Initialize MongoDB client
    mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        logger.Fatal("Failed to create MongoDB client: ", err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = mongoClient.Connect(ctx)
    if err != nil {
        logger.Fatal("Failed to connect to MongoDB: ", err)
    }
    defer mongoClient.Disconnect(ctx)

    mupVozilaClient := client.NewMupvozilaClient(&http.Client{}, "http://mupvozila_service:8081")
    authClient := client.NewAuthClient(&http.Client{}, "http://auth_service:8082")

    // Initialize repository
    repo := data.NewSaobracajnaPolicijaRepo(ctx, logger, mupVozilaClient, authClient)

    // Initialize handler
    handler := handler.NewSaobracajnaPolicijaHandler(repo, logger) // Ensure this function exists

    // Initialize router
    router := mux.NewRouter()

    // Routes for handling traffic accidents
    router.HandleFunc("/nesreca/new", handler.CreateNesreca).Methods(http.MethodPost)

    // Routes for handling violations
    router.HandleFunc("/prekrsaj/new", handler.CreatePrekrsaj).Methods(http.MethodPost)

    router.HandleFunc("/nesreca", handler.GetNesrece).Methods(http.MethodGet)
    router.HandleFunc("/prekrsaji", handler.GetPrekrsaji).Methods(http.MethodGet)
    router.HandleFunc("/nesreca/{id}", handler.DeleteNesreca).Methods(http.MethodDelete)
    router.HandleFunc("/prekrsaj/{id}", handler.DeletePrekrsaj).Methods(http.MethodDelete)

    router.HandleFunc("/cars", handler.GetAllCars).Methods(http.MethodGet)
    router.HandleFunc("/licenses/user/{jmbg}", handler.GetLicensesByUserJMBG).Methods(http.MethodGet)
    router.HandleFunc("/cars/plate/{license_plate}", handler.GetCarByLicensePlate).Methods(http.MethodGet)
    router.HandleFunc("/users/jmbg/{jmbg}", handler.GetUserByJMBG).Methods(http.MethodGet)
    headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"http://localhost:4200"}) // Update with your Angular app's origin

    handlerWithCORS := handlers.CORS(headers, methods, origins)(router)

    http.Handle("/", handlerWithCORS)
    server := &http.Server{
        Addr:         ":" + port,
        Handler:      handlerWithCORS,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    // Start the HTTP server in a separate goroutine
    go func() {
        logger.Println("Starting server on port ", port)
        err := server.ListenAndServe()
        if err != nil {
            logger.Fatalf("Server failed to start: %v", err)
        }
    }()

    // Wait for termination signal
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    sig := <-sigChan
    logger.Println("Received termination signal: ", sig)

    // Shutdown the server with timeout
    timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    server.Shutdown(timeoutCtx)
    logger.Println("Server stopped")
}
