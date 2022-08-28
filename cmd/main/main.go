package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/A-Victory/go-books/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.ResgisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
