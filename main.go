package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnka9755/go-03STUDENTAPP/internal/user"
	"github.com/jnka9755/go-03STUDENTAPP/package/boostrap"
	"github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()

	_ = godotenv.Load()

	log := boostrap.InitLooger()

	db, err := boostrap.DBConnection()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewRepository(log, db)
	userBusiness := user.NewBusiness(log, userRepository)
	userController := user.MakeEndpoints(userBusiness)

	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userController.Get).Methods("GET")
	router.HandleFunc("/users", userController.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userController.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")

	server := http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout!"),
		Addr:         "127.0.0.1:5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	error := server.ListenAndServe()

	if err != nil {
		log.Fatal(error)
	}
}
