package main

import (
	"github.com/gorilla/mux"
	"github.com/lipaysamart/withGoBookStore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))
}
