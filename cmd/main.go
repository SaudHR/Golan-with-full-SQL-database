package main

import (
	"Data/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegsiterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))

}
