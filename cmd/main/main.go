package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guoqiang/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// create the server
// tell golang where the server is

func main() {
	r := mux.NewRouter()
	// pass r to RegisterBookStoreRoutes and handle the CRUD
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
