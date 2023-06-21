package course

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jnka9755/go-03STUDENTAPP/package/meta"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
	}

	CreateReq struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

func MakeEndpoints(b Business) Endpoints {

	return Endpoints{
		Create: makeCreateEndpoint(b),
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

		if req.Name == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "name is required"})
			return
		}

		if req.StartDate == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "start_date is required"})
			return
		}

		if req.EndDate == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "end_date is required"})
			return
		}

		startDateParsed, _ := time.Parse("2006-01-02", req.StartDate)
		endDateParsed, _ := time.Parse("2006-01-02", req.EndDate)

		course := Course{
			Name:      req.Name,
			StartDate: startDateParsed,
			EndDate:   endDateParsed,
		}

		responseCourse, err := b.Create(&course)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 201, Data: responseCourse})
	}
}
