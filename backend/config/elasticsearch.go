package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

// ConnectElasticSearch ...
func (gc *GlobalConfig) ConnectElasticSearch() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			gc.Env.ElasticSearchEnv.ClientURL,
		},
		Username: gc.Env.ElasticSearchEnv.Username,
		Password: gc.Env.ElasticSearchEnv.Password,
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalln(err)
	}

	gc.ES = es

	log.Printf("es		| connected successfully\n")

}
