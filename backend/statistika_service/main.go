package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"statistika_service/client"
	"statistika_service/handler"
	"time"
)

func main() {
	dburi := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(dburi)

	dbClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8085"
	}

	saobracajnaPolicijaClient := client.NewSaobracajnaPolicijaClient("http://saobracajna_policija:8084")
	mupVozilaClient := client.NewMupVozilaClient("http://mupvozila_service:8081")
	lawCourtClient := client.NewLawCourtClient("http://law_court:8083")

	r := mux.NewRouter()

	r.HandleFunc("/statistikaPrekrsaja", handler.CreatePrekrsajiStatistika(saobracajnaPolicijaClient)).Methods("GET")
	r.HandleFunc("/statistikaNesreca", handler.CreateNesrecaStatistika(saobracajnaPolicijaClient)).Methods("GET")
	r.HandleFunc("/statistikaVozackihDozvola", handler.CreateVozackihDozvolaStatistika(mupVozilaClient)).Methods("GET")
	r.HandleFunc("/statistikaRegistrovanihVozila", handler.CreateRegistrovanihVozilaStatistika(mupVozilaClient)).Methods("GET")
	r.HandleFunc("/statistikaAuta", handler.CreateStatistikaAuta(mupVozilaClient)).Methods("GET")
	r.HandleFunc("/statistikaNalogaZaPretres", handler.CreateNaloziZaPretresStatistika(lawCourtClient)).Methods("GET")
	r.HandleFunc("/statistikaSaslusanja", handler.CreateStatistikaSaslusanja(lawCourtClient)).Methods("GET")
	r.HandleFunc("/statistikaPravnogZahteva", handler.CreateStatistikaPravnogZahteva(lawCourtClient)).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{
		Handler:      handlers.CORS(headers, methods, origins)(r),
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe error: %v\n", err)
		}
	}()

	log.Printf("Server running on port %s\n", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Println("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	dbClient.Disconnect(ctx)
}
