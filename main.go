package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MehulAgarwal07/to-do-go/controller"
	"github.com/MehulAgarwal07/to-do-go/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:5000", mux))

}
