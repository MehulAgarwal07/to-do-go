package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MehulAgarwal07/to-do-go/views"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	return mux
}
