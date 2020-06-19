package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MehulAgarwal07/to-do-go/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			// w.WriteHeader(201) //to change the response code
			json.NewEncoder(w).Encode(data)
		}
	}
}
