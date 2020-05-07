package models

import (
	"context"
	"time"

	"github.com/zerefwayne/that-meme/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Meme ...
type Meme struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	FileURL     string    `json:"file_url" bson:"file_url"`
	Name        string    `json:"name" bson:"name"`
	Tags        []string  `json:"tags" bson:"tags"`
	Description string    `json:"description" bson:"description"`
	Text        string    `json:"text" bson:"text"`
	Origin      string    `json:"origin" bson:"origin"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

// InsertMeme ...
func InsertMeme(m *Meme) error {

	result, err := config.Config.DB.Database("thatmemedev").Collection("memes").InsertOne(context.Background(), *m)

	if err != nil {
		return err
	}

	m.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return nil

}
