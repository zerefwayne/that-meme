package config

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// ConnectCache ...
func (c *Connections) ConnectCache() {

	redisURI := "redis://localhost"

	conn, err := redis.DialURL(redisURI)

	if err != nil {
		log.Fatal(err)
	}

	c.Cache = conn

	log.Printf("cache	| connected successfully: %s\n", redisURI)
}
