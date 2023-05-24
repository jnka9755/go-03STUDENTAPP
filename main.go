package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses).Methods("GET")

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

func getUsers(w http.ResponseWriter, r *http.Request) {

	time.Sleep(10 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getCourses(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
