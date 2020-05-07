package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello world!")

}

func main() {

	var connections Connections

	connections.ConnectDatabase()
	connections.ConnectCache()

	defer connections.Cache.Close()

	// ROUTER SETUP

	r := mux.NewRouter()

	r.HandleFunc("/", defaultHandler)

	// SERVER SETUP

	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":5000", handler)

}
