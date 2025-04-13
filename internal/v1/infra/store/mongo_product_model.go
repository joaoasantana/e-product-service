package store

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type MongoProduct struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description,omitempty"`
	CategoryID  bson.ObjectID `bson:"category_id"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}
