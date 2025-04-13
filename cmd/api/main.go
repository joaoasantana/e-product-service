package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/configs"
	"github.com/joaoasantana/e-product-service/internal/v1/api/http/category"
	"github.com/joaoasantana/e-product-service/internal/v1/api/http/product"
	"github.com/joaoasantana/e-product-service/pkg/util/connect"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

func main() {
	config := configs.LoadNewConfig()

	mongoDatabase := connect.MongoDatabase(config.Database)
	defer disconnectMongoDatabase(mongoDatabase)

	apiServer := gin.Default()

	group := apiServer.Group("/api/v1")
	{
		category.NewRouter(group, mongoDatabase)
		product.NewRouter(group, mongoDatabase)
	}

	if err := apiServer.Run(config.Server.Port); err != nil {
		disconnectMongoDatabase(mongoDatabase)
		panic(err)
	}
}

func disconnectMongoDatabase(database *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := database.Client().Disconnect(ctx); err != nil {
		panic(err)
	}
}
