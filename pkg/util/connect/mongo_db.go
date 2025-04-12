package connect

import (
	"context"
	"fmt"
	"github.com/joaoasantana/e-product-service/pkg/util/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"time"
)

func MongoDatabase(config configs.Database) *mongo.Database {
	opts := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s",
		config.Host, config.Port,
	))

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client.Database(config.Name)
}
