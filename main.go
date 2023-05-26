package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnka9755/go-03STUDENTAPP/internal/user"
)

func main() {

	router := mux.NewRouter()

	userBusiness := user.NewBusiness()
	userController := user.MakeEndpoints(userBusiness)

	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users", userController.Get).Methods("GET")
	router.HandleFunc("/users", userController.GetAll).Methods("GET")
	router.HandleFunc("/users", userController.Update).Methods("PATCH")
	router.HandleFunc("/users", userController.Delete).Methods("DELETE")

	server := http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout!"),
		Addr:         "127.0.0.1:5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
