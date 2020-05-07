package config

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// ConnectCache ...
func (gc *GlobalConfig) ConnectCache() {

	redisURI := "redis://localhost"

	conn, err := redis.DialURL(redisURI)

	if err != nil {
		log.Fatal(err)
	}

	gc.Cache = conn

	log.Printf("cache	| connected successfully: %s\n", redisURI)
}
