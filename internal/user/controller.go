package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jnka9755/go-03STUDENTAPP/package/meta"
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
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json: "meta,omitempty"`
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

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "invalid request format"})
			return
		}

		if req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "first name is required"})
			return
		}

		if req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "last name is required"})
			return
		}

		user := User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		}

		responseUser, err := b.Create(&user)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 201, Data: responseUser})
	}
}

func makeGetEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]

		user, err := b.Get(id)

		if err != nil {

			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: user})
	}
}

func makeGetAllEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		value := r.URL.Query()

		filters := Filters{
			FirstName: value.Get("first_name"),
			LastName:  value.Get("last_name"),
		}

		limit, _ := strconv.Atoi(value.Get("limit"))
		page, _ := strconv.Atoi(value.Get("page"))

		count, err := b.Count(filters)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}

		meta, err := meta.New(page, limit, count)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}

		users, err := b.GetAll(filters, meta.Offset(), meta.Limit())

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: users, Meta: meta})
	}
}

func makeUpdateEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var request UpdateReq

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Invalid request format"})
			return
		}

		if request.FirstName != nil && *request.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "first name is required"})
			return
		}

		if request.LastName != nil && *request.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "last name is required"})
			return
		}

		path := mux.Vars(r)
		id := path["id"]

		if err := b.Update(id, request.FirstName, request.LastName, request.Email, request.Phone); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "User doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "Successful update"})
	}
}

func makeDeleteEndpoint(b Business) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]

		if err := b.Delete(id); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "User doesn't exist"})
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "Successful delete"})
	}
}
