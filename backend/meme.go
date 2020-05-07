package main

import "time"

// Meme ...
type Meme struct {
	ID          string    `json:"id" bson:"_id"`
	FileURL     string    `json:"file_url" bson:"file_url"`
	Name        string    `json:"name" bson:"name"`
	Tags        []string  `json:"tags" bson:"tags"`
	Description string    `json:"description" bson:"description"`
	Text        string    `json:"text" bson:"text"`
	Origin      string    `json:"origin" bson:"origin"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
