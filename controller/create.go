package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MehulAgarwal07/to-do-go/model"
	"github.com/MehulAgarwal07/to-do-go/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
