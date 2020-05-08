package main

import (
	"log"
	"net/http"

	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/routes"
)

func init() {
	config.Config.LoadEnv()

	config.Config.ConnectDatabase()
	config.Config.ConnectS3()
	config.Config.ConnectElasticSearch()
}

func main() {

	handler := routes.NewHandler()

	log.Fatal(http.ListenAndServe(":5000", handler))

}
