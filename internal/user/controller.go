package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirstName string `json: "first_name"`
		LastName  string `json: "last_name"`
		Email     string `json: "email"`
		Phone     string `json: "phone"`
	}

	ErrorResponse struct {
		Error string `json: "error"`
	}
)

func MakeEndpoints(b Business) Endpoints {

	return Endpoints{
		Create: makeCreateEndpoint(b),
		Get:    makeGetEndpoint(b),
		GetAll: makeGetAllEndpoint(b),
		Update: makeUpdateEndpoint(b),
		Delete: makeDeleteEndpoint(b),
	}
}

func makeCreateEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var request CreateReq

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{"invalid request format"})
			return
		}

		if request.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{"first name is required"})
			return
		}

		if request.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{"last name is required"})
			return
		}

		user := User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Phone:     request.Phone,
		}

		responseUser, err := b.Create(&user)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
			return
		}

		json.NewEncoder(w).Encode(responseUser)
	}
}

func makeGetEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint getAll user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint update user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
