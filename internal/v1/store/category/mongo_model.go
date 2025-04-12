package category

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type MongoModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}
