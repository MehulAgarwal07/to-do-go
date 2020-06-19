package main

import (
	"net/http"

	"github.com/MehulAgarwal07/to-do-go/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)

}
