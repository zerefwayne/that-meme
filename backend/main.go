package main

import (
	"log"
	"net/http"

	"github.com/zerefwayne/that-meme/routes"
)

func main() {

	var connections Connections

	connections.ConnectDatabase()
	connections.ConnectCache()

	defer connections.Cache.Close()

	handler := routes.NewHandler()

	log.Fatal(http.ListenAndServe(":5000", handler))

}
