package main

import (
	"log"
	"net/http"

	"github.com/zerefwayne/that-meme/config"
	"github.com/zerefwayne/that-meme/routes"
)

func main() {

	config.Config.LoadEnv()

	config.Config.ConnectDatabase()

	config.Config.ConnectCache()
	defer config.Config.Cache.Close()

	config.Config.ConnectS3()

	config.Config.ConnectElasticSearch()

	handler := routes.NewHandler()

	log.Fatal(http.ListenAndServe(":5000", handler))

}
